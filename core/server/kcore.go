package server

import "github.com/gin-gonic/gin"

type KCore struct {
	*gin.Engine
}

func New() *KCore {
	return &KCore{
		gin.New(),
	}
}

func (this *KCore) Mount(class ...IClass) *KCore {
	for _, iClass := range class {
		iClass.Build(this)
	}
	return this
}
