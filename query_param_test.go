package learn_golang_web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameterExists(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000/hello?name=Nandes", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello Nandes", string(body))
}

func TestQueryParameterNotExists(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000/hello?name=", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello", string(body))
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000?first_name=nandes&last_name=simanjuntak", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	expect := "Hello nandes simanjuntak"
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	assert.Equal(t, expect, string(body))
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000?name=Fernandes&name=Ariadi&name=Simanjuntak", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	expect := "Fernandes Ariadi Simanjuntak"
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	assert.Equal(t, expect, string(body))

}
