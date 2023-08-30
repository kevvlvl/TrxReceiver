package transaction

import (
	"TrxReceiver/rdb"
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetTransaction(t *testing.T) {

	r, mock := redismock.NewClientMock()
	trx := trxTestSetup(r)

	stock := stockStub()
	stockBytes, _ := json.Marshal(stock)

	mock.ExpectGet(stock.IdStr()).SetVal(string(stockBytes[:]))

	resultStock, resultStockBytes := trx.GetTransaction(stock.IdStr())

	assert.Equal(t, stock.Id, resultStock.Id, "The result trx is different from the expected trx")
	assert.Equal(t, stockBytes, resultStockBytes, "The result trx bytes are different from the expected trx bytes")
}

func TestCreateTransaction(t *testing.T) {

	r, mock := redismock.NewClientMock()
	trx := trxTestSetup(r)

	stock := stockStub()
	stockBytes := stock.AsBytes()

	reqUrl := "localhost:4000/trx"
	req, _ := http.NewRequest("POST", reqUrl, bytes.NewBuffer([]byte(stockRequestBody())))

	mock.ExpectSet(stock.IdStr(), string(stockBytes[:]), 0).SetVal("ok")

	trx.CreateTransaction(req)
}

func TestUpdateTransaction(t *testing.T) {

	r, mock := redismock.NewClientMock()
	trx := trxTestSetup(r)

	stock := stockStub()
	stockBytes := stock.AsBytes()

	reqUrl := "localhost:4000/trx/1234"
	req, _ := http.NewRequest("PUT", reqUrl, bytes.NewBuffer([]byte(stockRequestBody())))

	mock.ExpectGet(stock.IdStr()).SetVal(string(stockBytes[:]))
	mock.ExpectSet(stock.IdStr(), string(stockBytes[:]), 0).SetVal("ok")

	trx.UpdateTransaction(req, stock.IdStr())
}

func TestWriteToRedis(t *testing.T) {

	r, mock := redismock.NewClientMock()
	trx := trxTestSetup(r)

	stock := stockStub()
	stockBytes := stock.AsBytes()

	mock.ExpectSet(stock.IdStr(), string(stockBytes[:]), 0).SetVal("ok")
	trx.writeToRedis(stock)
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

func stockRequestBody() string {
	return "{\"id\": 1234, \"symbol\": \"TEST\", \"name\": \"Test Stock\", \"Value\": 9001}"
}
