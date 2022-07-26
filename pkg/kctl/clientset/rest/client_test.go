package rest

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var cfg = &Config{
	Host:    "http://www.baidu.com",
	Timeout: time.Second * 10,
}

func TestNewRESTClient(t *testing.T) {
	req := NewRESTClient(cfg)
	rsp, err := req.R().Execute(http.MethodGet, "/")
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(string(rsp.Body()))
	}
}
