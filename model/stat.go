package model

type Stat struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    data   `json:"data"`
}
type data struct {
	Mid       int64 `json:"mid"`
	Following int64 `json:"following"`
	Whisper   int64 `json:"whisper"`
	Black     int   `json:"black"`
	Follower  int64 `json:"follower"`
}
