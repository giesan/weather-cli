package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type Error struct {
	Code    uint16 `json:"code"`
	Message string `json:"message"`
}

type ErrBody struct {
	Error `json:"error"`
}

func printAPIError(res *http.Response) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error(err.Error())
	}

	var errBody ErrBody
	err = json.Unmarshal(body, &errBody)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Error("Error returned from API", "errcode", errBody.Code, "errmessage", errBody.Message)
}
