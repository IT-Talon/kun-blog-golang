package main

import (
	"kun-blog-golang/core/server"
	"kun-blog-golang/pkg/kserver/ctls"
)

func main() {
	r := server.New()
	r.LoadHTMLGlob("./public/tpl/*")
	r.Mount("", ctls.NewPageCtl())
	r.Mount("v1", ctls.NewVersionCtl(), ctls.NewPostCtl(), ctls.NewPageCtl())
	r.Start()
}
