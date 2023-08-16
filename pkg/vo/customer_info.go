package vo

type CustomerInfo struct {
	Id      int64 `json:"id,omitempty"`
	Enabled bool  `json:"enabled"`
}
