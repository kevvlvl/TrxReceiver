package route

import (
	"TrxReceiver/rdb"
	"TrxReceiver/transaction"
	"context"
	"github.com/go-redis/redismock/v9"
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

	serveHttpTestRequest(w, r)

	assert.Equal(t, 200, w.Code, "The return code is not HTTP 200/OK")
	assert.Equal(t, "Root path", w.Body.String(), "The response body is not 'Root path'")
}

func TestHandleRouteGetHealth(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/health", nil)

	assert.Nil(t, err, "An error was returned for the new request")

	serveHttpTestRequest(w, r)

	assert.Equal(t, 200, w.Code, "The return code is not HTTP 200/OK")
	assert.Equal(t, "Up and Healthy", w.Body.String(), "The response body is not 'Root path'")
}

func TestHandleRoutePostStock(t *testing.T) {

	stock := stockStub()
	stockString := stock.AsString()

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/trx", strings.NewReader(stockString))
	assert.Nil(t, err, "An error was returned for the new request")

	redisMock := serveHttpTestRequest(w, r)

	redisMock.ExpectSet(stock.IdStr(), string(stock.AsBytes()[:]), 0).SetVal("ok")

	assert.Equal(t, 200, w.Code, "The return code is not HTTP 200/OK")
}

func serveHttpTestRequest(w *httptest.ResponseRecorder, r *http.Request) (redisMock redismock.ClientMock) {

	red, mock := redismock.NewClientMock()
	redisClientMock := rdb.RedisDB{
		Client:  red,
		Context: context.Background(),
	}

	chiRouter := Router(&redisClientMock)
	chiRouter.handleRoutes()
	chiRouter.Router.ServeHTTP(w, r)

	return mock
}

func stockStub() transaction.Stock {
	return transaction.Stock{
		Id:     1234,
		Symbol: "TEST",
		Name:   "Test Stock",
		Value:  9001,
	}
}
