package netops

import (
	"bytes"
	"encoding/json"
	"imageGO/internal/util"
	"net/http"
	"testing"
)

func TestFetchGet(t *testing.T) {
	url := "https://httpbin.org/get?name=go"

	req, err := http.NewRequest("GET", url, nil)

	util.CheckForFailureTest(err, t)

	res := Fetch(req)

	var parsed_response struct {
		Url string `json:"url"`
	}
	err = json.Unmarshal([]byte(res), &parsed_response)

	util.CheckForFailureTest(err, t)

	if parsed_response.Url != url {
		t.Fatal("GET FAILED: ", err)
	}
}

func TestFetchPost(t *testing.T) {
	url := "https://httpbin.org/post"
	body := "name=go"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	util.CheckForFailureTest(err, t)

	res := Fetch(req)

	var parsed_response struct {
		Data string `json:"data"`
	}

	err = json.Unmarshal([]byte(res), &parsed_response)
	util.CheckForFailureTest(err, t)

	if parsed_response.Data != body {
		t.Fatal("GET FAILED: Sent and recieved body are not same.")
	}
}
