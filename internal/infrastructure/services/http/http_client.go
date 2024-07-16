package http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"time"
)

type HttpClientSettings struct {
	Host    string
	Timeout string
}

type httpClient[Request any, Response any] struct {
	settings HttpClientSettings
}

type HttpResponse[Response any] struct {
	StatusCode int
	Body       Response
}

type HttpClientError struct {
	StatusCodeMessage int
	ErrorMessage      string
}

func (h *HttpClientError) Error() string {
	return h.ErrorMessage
}

func (h *HttpClientError) StatusCode() int {
	return h.StatusCodeMessage
}

func NewHttpClient[Request any, Response any](settings HttpClientSettings) HttpClientInterface[Request, Response] {
	return &httpClient[Request, Response]{settings}
}

func (httpClient *httpClient[Request, Response]) Post(ctx context.Context, url string, body Request) (*HttpResponse[Response], error) {
	var response Response
	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(body)
	if err != nil {
		msg := fmt.Sprintf("Error encoding a message, error: %v", err.Error())
		// TODO: Add logs
		return nil, returnHttpError(nil, msg)
	}

	endpoint := fmt.Sprintf("%s/%s", httpClient.settings.Host, url)
	request, errReq := http.NewRequest("POST", endpoint, &buffer)
	if errReq != nil {
		msg := fmt.Sprintf("Error creating the http post request, error: %v", errReq.Error())
		return nil, returnHttpError(nil, msg)
	}
	var timeout int
	if httpClient.settings.Timeout == "" {
		timeout = 0
	} else {
		timeout, _ = strconv.Atoi(httpClient.settings.Timeout)
	}
	client := &http.Client{Timeout: time.Duration(timeout) * time.Millisecond}
	res, errClient := client.Do(request)
	if errClient != nil {
		msg := fmt.Sprintf("Error while executing client Do for http post, error: %v", errClient.Error())
		return nil, returnHttpError(res, msg)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// TODO: Logging here
		}
	}(res.Body)

	statusOK := []int{http.StatusOK, http.StatusCreated}
	if !slices.Contains(statusOK, res.StatusCode) {
		var errResponse map[string]interface{}
		_ = json.NewDecoder(res.Body).Decode(&errResponse)
		msg := fmt.Sprintf("Status Code: %d with error: %#v", res.StatusCode, errResponse)
		return nil, returnHttpError(res, msg)
	}

	if res.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	errDecoding := json.NewDecoder(res.Body).Decode(&response)
	if errDecoding != nil {
		msg := fmt.Sprintf("Error parsing response with msg: %v", errDecoding)
		return nil, returnHttpError(res, msg)
	}

	return &HttpResponse[Response]{
		Body:       response,
		StatusCode: res.StatusCode,
	}, nil
}

func (httpClient *httpClient[Request, Response]) Put(ctx context.Context, url string, body Request) (*HttpResponse[Response], error) {
	return nil, nil
}

func returnHttpError(res *http.Response, msg string) *HttpClientError {

	statusCode := http.StatusInternalServerError
	if res != nil {
		statusCode = res.StatusCode
	}
	return &HttpClientError{
		StatusCodeMessage: statusCode,
		ErrorMessage:      msg,
	}
}
