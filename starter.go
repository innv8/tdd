package starter

import (
	"fmt"
	"math"
	"net/http"
)

func SayHello(name string) string {
	return fmt.Sprintf("Hello %v, Welcome!", name)
}

func OddOrEven(number int) string {
	criteria := math.Mod(float64(number), 2)
	if criteria == 1 || criteria == -1 {
		return fmt.Sprintf("%d is an odd number", number)
	}
	return fmt.Sprintf("%d is an even number", number)
}

//CheckHealth
// this is like a controller in a web-api
// when testing, we will leverage the net/http/httptest
func CheckHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "health check passed")
}
