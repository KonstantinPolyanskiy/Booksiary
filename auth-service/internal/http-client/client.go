package http_client

import "net/http"

type Client struct {
	*http.Client
}

func Default() Client {
	return Client{http.DefaultClient}
}
