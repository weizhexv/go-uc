package vo

import (
	"encoding/json"
	"time"
)

type TokenInfo struct {
	Uid       int64     `json:"uid,omitempty"`
	Domain    string    `json:"domain,omitempty"`
	Platform  string    `json:"platform,omitempty"`
	ExpiresAt time.Time `json:"expiresAt,omitempty"`
}

func (t *TokenInfo) isExpired() bool {
	return time.Now().After(t.ExpiresAt)
}

func NewTokenInfo(uid int64, domain string, platform string, expiresAt time.Time) *TokenInfo {
	return &TokenInfo{
		Uid:       uid,
		Domain:    domain,
		Platform:  platform,
		ExpiresAt: expiresAt,
	}
}

func (i *TokenInfo) String() string {
	if j, err := json.Marshal(i); err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
