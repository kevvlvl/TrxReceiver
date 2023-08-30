package transaction

import (
	"bytes"
	"encoding/json"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestGetTransaction(t *testing.T) {

	r, mock := redismock.NewClientMock()

	trx := stubTrx()
	trxBytes, _ := json.Marshal(trx)

	mock.ExpectGet(intToStr(trx.Id)).SetVal(string(trxBytes[:]))

	resultTrx, resultTrxBytes := GetTransaction(trx.Id)

	assert.Equal(t, trx.Id, resultTrx.Id, "The result trx is different from the expected trx")
	assert.Equal(t, trxBytes, resultTrxBytes, "The result trx bytes are different from the expected trx bytes")
}

func TestCreateTransaction(t *testing.T) {

	trx := stubTrx()
	trxBytes, _ := json.Marshal(trx)

	_, mock := redismock.NewClientMock()

	reqUrl := "localhost:4000/trx"
	req, _ := http.NewRequest("POST", reqUrl, bytes.NewBuffer([]byte(stubTrxAsStr())))

	mock.ExpectSet(intToStr(trx.Id), string(trxBytes[:]), 0).SetVal("ok")

	CreateTransaction(req)
}

func TestUpdateTransaction(t *testing.T) {

	trx := stubTrx()
	trxBytes, _ := json.Marshal(trx)

	_, mock := redismock.NewClientMock()

	reqUrl := "localhost:4000/trx/1234"
	req, _ := http.NewRequest("PUT", reqUrl, bytes.NewBuffer([]byte(stubTrxAsStr())))

	mock.ExpectGet(intToStr(trx.Id)).SetVal(string(trxBytes[:]))
	mock.ExpectSet(intToStr(trx.Id), string(trxBytes[:]), 0).SetVal("ok")

	UpdateTransaction(req, 1234)
}

func TestIntToStr(t *testing.T) {

	number := 50
	expectedStr := strconv.FormatInt(int64(number), IntBase)

	str := intToStr(50)

	assert.Equal(t, expectedStr, str, "The string equivalent does not represent the int")
}

func TestWriteToRedis(t *testing.T) {

	_, mock := redismock.NewClientMock()

	trx := stubTrx()

	trxJson, err := json.Marshal(trx)

	if err != nil {
		t.Error("Unexpected error marshalling test struct")
	}

	mock.ExpectSet(intToStr(trx.Id), string(trxJson[:]), 0).SetVal("ok")
	writeToRedis(trx)
}

func stubTrx() Stock {
	return Stock{
		Id:     1234,
		Symbol: "TEST",
		Name:   "Test Stock",
		Value:  9001,
	}
}

func stubTrxAsStr() string {
	return "{\"id\": 1234, \"symbol\": \"TEST\", \"name\": \"Test Stock\", \"Value\": 9001}"
}
