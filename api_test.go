package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

const API_URL = "http://localhost:8080"

type StatusCheck struct {
	Ping string `json:"ping"`
}

type StatusResponse struct {
	Pong string `json:"pong"`
}

func TestApiPing(t *testing.T) {
	ping, err := json.Marshal(StatusCheck{Ping: "ping!"})
	if err != nil {
		t.Error("Couldn't serialize the json.")
	}

	req, err2 := http.NewRequest("GET", API_URL+"/ping", bytes.NewBuffer(ping))
	if err2 != nil {
		t.Error("Couldn't build the http request")
	}

	client := &http.Client{}
	resp, err3 := client.Do(req)

	if err3 != nil {
		t.Error("Couldn't make the http request to given url")
	}

	defer resp.Body.Close()

	t.Logf("response status:%s\n", resp.Status)
	t.Logf("response headers: %s\n", resp.Header)

	if resp.StatusCode != 200 {
		t.Errorf("got error on status. got %d, want 200", resp.StatusCode)
	}

	var pongResp StatusResponse
	body, _ := io.ReadAll(resp.Body)
	err4 := json.Unmarshal(body, &pongResp)

	if err4 != nil {
		t.Errorf("Couldn't deserialize the json from response: %s", err4)
	}

	wantedResponse := "pong!"

	if pongResp.Pong != wantedResponse {
		t.Errorf("response is invalid, got %s, want %s", pongResp.Pong, wantedResponse)
	}

}
