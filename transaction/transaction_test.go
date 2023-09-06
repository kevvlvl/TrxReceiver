package transaction

import (
	"TrxReceiver/rdb"
	"context"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestGetTransaction(t *testing.T) {

	r, mock := redismock.NewClientMock()
	trx := trxTestSetup(r)

	stock := stockStub()

	mock.ExpectGet(stock.IdStr()).SetVal(stock.AsString())

	resultStockBytes := trx.GetTransaction(stock.IdStr())

	assert.Equal(t, stock.AsBytes(), resultStockBytes, "The result trx bytes are different from the expected trx bytes")
}

func TestCreateTransaction(t *testing.T) {

	r, mock := redismock.NewClientMock()
	trx := trxTestSetup(r)

	stock := stockStub()

	reqUrl := "localhost:4000/trx"
	req, _ := http.NewRequest("POST", reqUrl, strings.NewReader(stock.AsString()))

	mock.ExpectSet(stock.IdStr(), stock.AsString(), 0).SetVal("ok")

	trx.CreateTransaction(req)
}

func TestUpdateTransaction(t *testing.T) {

	r, mock := redismock.NewClientMock()
	trx := trxTestSetup(r)

	stock := stockStub()

	reqUrl := "localhost:4000/trx/1234"
	req, _ := http.NewRequest("PUT", reqUrl, strings.NewReader(stock.AsString()))

	mock.ExpectGet(stock.IdStr()).SetVal(stock.AsString())
	mock.ExpectSet(stock.IdStr(), stock.AsString(), 0).SetVal("ok")

	trx.UpdateTransaction(req, stock.IdStr())
}

func TestWriteToRedis(t *testing.T) {

	r, mock := redismock.NewClientMock()
	trx := trxTestSetup(r)

	stock := stockStub()

	mock.ExpectSet(stock.IdStr(), stock.AsString(), 0).SetVal("ok")
	trx.writeToRedis(&stock)
}

func trxTestSetup(r *redis.Client) Trx {

	return Trx{
		Redis: &rdb.RedisDB{
			Client:  r,
			Context: context.Background(),
		},
	}
}

func stockStub() Stock {
	return Stock{
		Id:     1234,
		Symbol: "TEST",
		Name:   "Test Stock",
		Value:  9001,
	}
}
