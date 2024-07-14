package http_client

import "context"

type HttpClientInterface[Request any, Response any] interface {
	Post(ctx context.Context, url string, body Request) (*HttpResponse[Response], error)
	Put(ctx context.Context, url string, body Request) (*HttpResponse[Response], error)
}
