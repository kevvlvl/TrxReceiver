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
	"strconv"
	"testing"
)

func TestGetTransaction(t *testing.T) {

	r, mock := redismock.NewClientMock()
	trx := TrxTestSetup(r)

	stock := stubStock()
	stockBytes, _ := json.Marshal(stock)

	mock.ExpectGet(intToStr(stock.Id)).SetVal(string(stockBytes[:]))

	resultStock, resultStockBytes := trx.GetTransaction(stock.Id)

	assert.Equal(t, stock.Id, resultStock.Id, "The result trx is different from the expected trx")
	assert.Equal(t, stockBytes, resultStockBytes, "The result trx bytes are different from the expected trx bytes")
}

func TestCreateTransaction(t *testing.T) {

	r, mock := redismock.NewClientMock()
	trx := TrxTestSetup(r)

	stock := stubStock()
	stockBytes, _ := json.Marshal(stock)

	reqUrl := "localhost:4000/trx"
	req, _ := http.NewRequest("POST", reqUrl, bytes.NewBuffer([]byte(stubStockAsStr())))

	mock.ExpectSet(intToStr(stock.Id), string(stockBytes[:]), 0).SetVal("ok")

	trx.CreateTransaction(req)
}

func TestUpdateTransaction(t *testing.T) {

	r, mock := redismock.NewClientMock()
	trx := TrxTestSetup(r)

	stock := stubStock()
	stockBytes, _ := json.Marshal(stock)

	reqUrl := "localhost:4000/trx/1234"
	req, _ := http.NewRequest("PUT", reqUrl, bytes.NewBuffer([]byte(stubStockAsStr())))

	mock.ExpectGet(intToStr(stock.Id)).SetVal(string(stockBytes[:]))
	mock.ExpectSet(intToStr(stock.Id), string(stockBytes[:]), 0).SetVal("ok")

	trx.UpdateTransaction(req, 1234)
}

func TestIntToStr(t *testing.T) {

	number := 50
	expectedStr := strconv.FormatInt(int64(number), IntBase)

	str := intToStr(50)

	assert.Equal(t, expectedStr, str, "The string equivalent does not represent the int")
}

func TestWriteToRedis(t *testing.T) {

	r, mock := redismock.NewClientMock()
	trx := TrxTestSetup(r)

	stock := stubStock()
	stockJson, err := json.Marshal(stock)

	if err != nil {
		t.Error("Unexpected error marshalling test struct")
	}

	mock.ExpectSet(intToStr(stock.Id), string(stockJson[:]), 0).SetVal("ok")
	trx.writeToRedis(stock)
}

func TrxTestSetup(r *redis.Client) Trx {

	return Trx{
		Redis: &rdb.RedisDB{
			Client:  r,
			Context: context.Background(),
		},
	}
}

func stubStock() Stock {
	return Stock{
		Id:     1234,
		Symbol: "TEST",
		Name:   "Test Stock",
		Value:  9001,
	}
}

func stubStockAsStr() string {
	return "{\"id\": 1234, \"symbol\": \"TEST\", \"name\": \"Test Stock\", \"Value\": 9001}"
}
