package starter_test

import (
	starter "github.com/innv8/tdd"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("Sam")
	assert.Equal(t, "Hello Sam, Welcome!", greeting)
}

//func TestOddOrEven(t *testing.T) {
//	assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
//	assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
//	assert.Equal(t, "43 is an odd number", starter.OddOrEven(43))
//}

func TestOddOrEven(t *testing.T) {
	// to test cases e.g happy path
	t.Run("Check non-negative numbers", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		assert.Equal(t, "43 is an odd number", starter.OddOrEven(43))
	})

	t.Run("Check negative numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
	})

}

func TestCheckHealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "http://mysite.com/example", nil)
		w := httptest.NewRecorder()
		starter.CheckHealth(w, r)
		response := w.Result()
		body, err := io.ReadAll(response.Body)
		assert.Equal(t, "health check passed", string(body))
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, "text/plain; charset=utf-8", response.Header.Get("content-type"))
		assert.Equal(t, nil, err)

	})
}
