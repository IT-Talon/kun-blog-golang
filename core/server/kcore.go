package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type KCore struct {
	g *gin.RouterGroup
	*gin.Engine
}

func New() *KCore {
	return &KCore{
		Engine: gin.Default(),
	}
}

func (this *KCore) Mount(group string, class ...IClass) *KCore {
	this.g = this.Group(group)
	for _, iClass := range class {
		iClass.Build(this)
	}
	return this
}

func (this *KCore) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *KCore {
	this.g.Handle(httpMethod, relativePath, handlers...)
	return this
}

func (this *KCore) Start() *KCore {
	port := 80
	this.Run(fmt.Sprintf(":%d", port))
	return this
}
