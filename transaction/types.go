package transaction

import "TrxReceiver/rdb"

type Stock struct {
	Id            int `json:"id"`
	idStr         string
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	Value         float32 `json:"value"`
	stockBytes    []byte
	stockBytesStr string
}

type Trx struct {
	Redis *rdb.RedisDB
}
