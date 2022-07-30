package model

type UpStat struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	TTL     int        `json:"ttl"`
	Data    UpStatData `json:"data"`
}
type Archive struct {
	View int64 `json:"view"`
}
type Article struct {
	View int64 `json:"view"`
}
type UpStatData struct {
	Archive Archive `json:"archive"`
	Article Article `json:"article"`
	Likes   int64   `json:"likes"`
}
