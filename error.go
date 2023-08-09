package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Error struct {
	Code    uint16 `json:"code"`
	Massage string `json:"message"`
}

type ErrBody struct {
	Error `json:"error"`
}

func printAPIError(res *http.Response) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var errBody ErrBody
	err = json.Unmarshal(body, &errBody)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Error code: %d\nError message: %s\n", errBody.Code, errBody.Massage)
}
