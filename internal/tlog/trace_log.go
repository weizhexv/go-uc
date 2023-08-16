package tlog

import (
	"dghire.com/libs/go-common/constants"
	dgctx "dghire.com/libs/go-common/context"
	dgsys "dghire.com/libs/go-common/sys"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go-uc/pkg/var/domains"
	"os"
	"sort"
)

type Ctx struct {
	*logrus.Entry
	*dgctx.DgContext
	dm domains.Domain
}

var tlog = initTLog()

func (c *Ctx) Domain() domains.Domain {
	return c.dm
}

func NewCtx(c *dgctx.DgContext) *Ctx {
	if len(c.TraceId) == 0 {
		c.TraceId = uuid.NewString()
	}
	return &Ctx{
		Entry:     newEntry(c.TraceId),
		DgContext: c,
		dm:        domains.Parse(c.Domain),
	}
}

func initTLog() *logrus.Logger {
	level, err := logrus.ParseLevel("INFO")
	if err != nil {
		panic(err)
	}

	return &logrus.Logger{
		Out:       os.Stdout,
		Formatter: setFormatter(),
		Level:     level,
	}
}

func setFormatter() logrus.Formatter {
	return &logrus.TextFormatter{
		DisableQuote:           true,
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05.999999",
		DisableSorting:         true,
		DisableLevelTruncation: true,
		PadLevelText:           false,
		SortingFunc: func(strings []string) {
			sort.Slice(strings, func(i, j int) bool {
				if strings[i] == "level" {
					return true
				}
				return false
			})
		},
	}
}

func newEntry(traceId string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		constants.TraceId: traceId,
		"goid":            dgsys.QuickGetGoRoutineId(),
	})
}
