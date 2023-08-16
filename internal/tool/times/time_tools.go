package times

import (
	"github.com/rolandhe/daog/ttypes"
	"time"
)

func NowDatetime() ttypes.NormalDatetime {
	return ttypes.NormalDatetime(time.Now())
}

func NowDate() ttypes.NormalDate {
	return ttypes.NormalDate(time.Now())
}

func NowNilableDatetime() ttypes.NilableDatetime {
	return *ttypes.FromDatetime(time.Now())
}

func NowNilableDate() *ttypes.NilableDate {
	return ttypes.FromDate(time.Now())
}

func ParseNilableDate(value string) (*ttypes.NilableDate, error) {
	if len(value) == 0 {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return nil, err
	}
	return ttypes.FromDate(t), nil
}
