package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	link := "https://api.weeek.net/public/v1/tm/tasks?day=31.03.2024&projectId=6&all=1"
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, link, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", getWeeekToken())
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("response status code:", res.StatusCode)

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	fmt.Println(res.Header.Get("Content-Type"))
	fmt.Println(readResponse(res.Body))
}

func getWeeekToken() string {
	m, err := godotenv.Read(".env")
	if err != nil {
		panic(err)
	}
	token := m["WEEEK_API_TOKEN"]
	token = "Bearer " + token
	fmt.Printf("got weeek token: %s\n", token)
	return token
}

func readResponse(r io.Reader) string {
	buf := make([]byte, 2048)
	bytes := []byte{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			bytes = append(bytes, b)
		}
		if err == io.EOF {
			var out string = string(bytes)
			return out
		}
		if err != nil {
			panic(err)
		}
	}
}
