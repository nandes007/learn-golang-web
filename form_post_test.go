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

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	first_name := request.PostForm.Get("first_name")
	last_name := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", first_name, last_name)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Nandes&last_name=Simanjuntak")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:9000", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	expected := "Hello Nandes Simanjuntak"

	assert.Equal(t, expected, string(body))
}