package clientset

import (
	"fmt"
	"kun-blog-golang/pkg/kctl/clientset/rest"
	"testing"
	"time"
)

var cfg = &rest.Config{
	Host:    "http://127.0.0.1",
	Timeout: time.Second * 10,
}

func TestNewClientSetForConfig(t *testing.T) {
	client := NewClientSetForConfig(cfg)
	ver, err := client.V1().Version().Get()
	fmt.Println(ver, err)
}
