package transaction

import (
	"encoding/json"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGetTransaction(t *testing.T) {

}

func TestParseTransactionBody(t *testing.T) {

}

func TestIntToStr(t *testing.T) {

	number := 50
	expectedStr := strconv.FormatInt(int64(number), IntBase)

	str := intToStr(50)

	assert.Equal(t, expectedStr, str, "The string equivalent does not represent the int")
}

func TestWriteToRedis(t *testing.T) {

	db, mock := redismock.NewClientMock()

	trx := Transaction{
		Id:     1234,
		Symbol: "TEST",
		Name:   "Test Transaction",
		Value:  9001,
	}

	trxJson, err := json.Marshal(trx)

	if err != nil {
		t.Error("Unexpected error marshalling test struct")
	}

	mock.ExpectSet(intToStr(trx.Id), string(trxJson[:]), 0)
	writeToRedis(trx, db)
}
