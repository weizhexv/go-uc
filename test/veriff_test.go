package test

import (
	dgctx "dghire.com/libs/go-common/context"
	dghttp "dghire.com/libs/go-httpclient"
	dglogger "dghire.com/libs/go-logger"
	"github.com/google/uuid"
	"go-uc/internal/proxy/veriff"
	"testing"
)

func TestIsVerified(t *testing.T) {
	dghttp.GlobalHttpClient.UseMonitor = false
	ctx := &dgctx.DgContext{TraceId: uuid.NewString()}
	verified, err := veriff.IsVerified(ctx, "863a61cf-9f6d-4047-b297-c3eba8df2a0f")
	if err != nil {
		dglogger.Errorln(ctx, err)
		return
	}
	dglogger.Infoln(ctx, verified)
}
