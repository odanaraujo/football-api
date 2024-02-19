package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log/slog"
	"net/http"
	"os"
)

const (
	URL            = "API"
	METHOD         = "GET"
	AUTHENTICATION = "AUTHENTICATION"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("error loading .env file", err)
		return
	}

	url := os.Getenv(URL)
	authentication := os.Getenv(AUTHENTICATION)
	client := http.Client{}

	req, err := http.NewRequest(METHOD, url, nil)

	if err != nil {
		slog.Error("error request client", err)
		return
	}

	req.Header.Add("Authorization", authentication)

	res, err := client.Do(req)

	if err != nil {
		slog.Error("error do client", err)
		return
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		slog.Error("error read body", err)
		return
	}

	fmt.Println(string(body))

}
