package route

import (
	"TrxReceiver/rdb"
	"TrxReceiver/transaction"
	"context"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleRouteGetRoot(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)

	assert.Nil(t, err, "An error was returned for the new request")

	redisDb, _ := redisMock()
	serveHttpTestRequest(w, r, redisDb)

	assert.Equal(t, 200, w.Code, "The return code is not HTTP 200/OK")
	assert.Equal(t, "Root path", w.Body.String(), "The response body is not 'Root path'")
}

func TestHandleRouteGetHealth(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/health", nil)

	assert.Nil(t, err, "An error was returned for the new request")

	redisDb, _ := redisMock()
	serveHttpTestRequest(w, r, redisDb)

	assert.Equal(t, 200, w.Code, "The return code is not HTTP 200/OK")
	assert.Equal(t, "Up and Healthy", w.Body.String(), "The response body is not 'Root path'")
}

func TestHandleRoutePostStock(t *testing.T) {

	stock := stockStub()
	stockString := stock.AsString()

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/trx", strings.NewReader(stockString))
	assert.Nil(t, err, "An error was returned for the new request")

	redisDb, mock := redisMock()

	mock.ExpectSet(stock.IdStr(), stock.AsString(), 0).SetVal("ok")

	serveHttpTestRequest(w, r, redisDb)

	assert.Equal(t, 200, w.Code, "The return code is not HTTP 200/OK")
}

func redisMock() (client *redis.Client, mock redismock.ClientMock) {
	return redismock.NewClientMock()
}

func serveHttpTestRequest(w *httptest.ResponseRecorder, r *http.Request, client *redis.Client) {

	redisClientMock := rdb.RedisDB{
		Client:  client,
		Context: context.Background(),
	}

	chiRouter := Router(&redisClientMock)
	chiRouter.handleRoutes()
	chiRouter.Router.ServeHTTP(w, r)
}

func stockStub() transaction.Stock {
	return transaction.Stock{
		Id:     1234,
		Symbol: "TEST",
		Name:   "Test Stock",
		Value:  9001,
	}
}
