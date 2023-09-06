package transaction

import "TrxReceiver/rdb"

type Stock struct {
	Id            int64 `json:"id"`
	idStr         string
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	Value         float32 `json:"value"`
	stockBytes    []byte
	stockBytesStr string
}

type Option struct {
	Id             int64 `json:"id"`
	idStr          string
	Stock          Stock `json:"stock"`
	Type           OptionType
	optionBytes    []byte
	optionBytesStr string
}

type FinancialInstrument interface {
	AsBytes() []byte
	AsString() string
	IdStr() string
}

type OptionType string

const (
	OptionCall = "Call"
	OptionPut  = "Put"
)

type Trx struct {
	Redis *rdb.RedisDB
}
