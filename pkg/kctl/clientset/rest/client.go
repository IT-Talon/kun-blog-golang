package rest

import (
	"github.com/go-resty/resty/v2"
)

type RESTClient struct {
	*resty.Client
}

func NewRESTClient(cfg *Config) *RESTClient {
	rc := resty.New()
	rc.SetBaseURL(cfg.Host)
	rc.SetTimeout(cfg.Timeout)
	return &RESTClient{
		Client: rc,
	}
}
