package veriff

import (
	"crypto/hmac"
	"crypto/sha256"
	dgctx "dghire.com/libs/go-common/context"
	dghttp "dghire.com/libs/go-httpclient"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-uc/internal/dal"
	"go-uc/vconfig"
)

const veriffSuccessCode = 9001

func IsIdCardMatched(ctx *dgctx.DgContext, sessionId string, emp *dal.Employee) (bool, error) {
	headers := buildHeaders(sessionId)
	url := fmt.Sprintf(vconfig.Veriff()["person-url"], sessionId)
	resp, err := dghttp.GlobalHttpClient.DoGet(ctx, url, map[string]string{}, headers)
	if err != nil {
		return false, err
	}

	mp := map[string]any{}
	json.Unmarshal(resp, &mp)
	if mp["person"] == nil {
		return false, nil
	}

	personMP, ok := mp["person"].(map[string]any)
	if !ok {
		return false, nil
	}

	return personMP["idCode"] == emp.IdCard.StringNilAsEmpty() || personMP["idCode"] == emp.Passport.StringNilAsEmpty(), nil
}

func IsVerified(ctx *dgctx.DgContext, sessionId string) (bool, error) {
	headers := buildHeaders(sessionId)
	url := fmt.Sprintf(vconfig.Veriff()["decision-url"], sessionId)
	resp, err := dghttp.GlobalHttpClient.DoGet(ctx, url, map[string]string{}, headers)
	if err != nil {
		return false, err
	}

	mp := map[string]any{}
	json.Unmarshal(resp, &mp)
	if mp["verification"] == nil {
		return false, nil
	}

	vMP, ok := mp["verification"].(map[string]any)
	if !ok {
		return false, nil
	}

	return vMP["code"] == veriffSuccessCode, nil
}

func buildHeaders(sessionId string) map[string]string {
	return map[string]string{
		"Content-Type":     "application/json",
		"X-HMAC-SIGNATURE": HmacSHA256(sessionId),
		"X-AUTH-CLIENT":    vconfig.Veriff()["public-key"],
	}
}

func HmacSHA256(s string) string {
	mac := hmac.New(sha256.New, []byte(vconfig.Veriff()["private-key"]))
	mac.Write([]byte(s))
	return hex.EncodeToString(mac.Sum(nil))
}
