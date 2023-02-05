package learn_golang_web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(400)
		fmt.Fprint(writer, "name is empty")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest("GET", "http:/localhost:9000", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	statusCode := response.StatusCode
	statusResponse := response.Status

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

	assert.Equal(t, 400, statusCode)
	assert.Equal(t, "400 Bad Request", statusResponse)
	assert.Equal(t, "name is empty", string(body))
}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest("GET", "http:/localhost:9000?name=Nandes", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	statusCode := response.StatusCode
	statusResponse := response.Status

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

	assert.Equal(t, 200, statusCode)
	assert.Equal(t, "200 OK", statusResponse)
	assert.Equal(t, "Hello Nandes", string(body))
}
