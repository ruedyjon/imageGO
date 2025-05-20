package netops

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestFetchGet(t *testing.T) {
	url := "https://httpbin.org/get?name=go"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal("GET FAILED: ", err)
	}

	res, err := Fetch(req)
	if err != nil {
		t.Fatal("GET FAILED: ", err)
	}

	var parsed_response struct {
		Url string `json:"url"`
	}
	err = json.Unmarshal([]byte(res), &parsed_response)
	if err != nil {
		t.Fatal("GET FAILED: ", err)
	}

	if parsed_response.Url != url {
		t.Fatal("GET FAILED: ", err)
	}
}

func TestFetchPost(t *testing.T) {
	url := "https://httpbin.org/post"
	body := "name=go"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Fatal("GET FAILED: ", err)
	}

	res, err := Fetch(req)
	if err != nil {
		t.Fatal("POST FAILED: ", err)
	}
	var parsed_response struct {
		Data string `json:"data"`
	}

	err = json.Unmarshal([]byte(res), &parsed_response)
	if err != nil {
		t.Fatal("GET FAILED: ", err)
	}

	if parsed_response.Data != body {
		t.Fatal("POST FAILED: Sent and recieved body are not same.")
	}
}
