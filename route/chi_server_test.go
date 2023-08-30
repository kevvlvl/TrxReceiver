package route

import (
	"TrxReceiver/rdb"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRouteGetRoot(t *testing.T) {

	redisMock, _ := redismock.NewClientMock()
	redisClientMock := rdb.RedisDB{
		Client: redisMock,
	}

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)

	assert.Nil(t, err, "An error was returned for the new request")

	Router(&redisClientMock).Router.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code, "The return code is not HTTP 200/OK")
	assert.Equal(t, "Root path", w.Body.String(), "The response body is not 'Root path'")

}

func TestHandleRouteGetHealth(t *testing.T) {

}

func TestParseTrxId(t *testing.T) {

}
