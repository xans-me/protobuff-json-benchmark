package benchmarks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	jsonHttp "github.com/xans-me/protobuff-json-benchmark/json-http"
)

func init() {
	go jsonHttp.Start()
	time.Sleep(time.Second)
}

func BenchmarkJSONHTTP(b *testing.B) {
	client := &http.Client{}

	for n := 0; n < b.N; n++ {
		doPost(client, b)
	}
}

func doPost(client *http.Client, b *testing.B) {
	u := &jsonHttp.User{
		Email:    "mulia.ichsan@gmail.com",
		Name:     "Mulia Ichsan",
		Password: "thisispassword00",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(u)

	resp, err := client.Post("http://127.0.0.1:60001/", "application/json", buf)
	if err != nil {
		b.Fatalf("http request failed: %v", err)
	}

	defer resp.Body.Close()

	// We need to parse response to have a fair comparison as gRPC does it
	var target jsonHttp.Response
	decodeErr := json.NewDecoder(resp.Body).Decode(&target)
	if decodeErr != nil {
		b.Fatalf("unable to decode json: %v", decodeErr)
	}

	if target.Code != 200 || target.User == nil || target.User.ID != "1000000" {
		b.Fatalf("http response is wrong: %v", resp)
	}
}
