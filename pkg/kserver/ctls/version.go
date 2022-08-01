package ctls

import (
	"github.com/gin-gonic/gin"
	"kun-blog-golang/core/server"
	v1 "kun-blog-golang/pkg/apis/v1"
	"kun-blog-golang/pkg/kctl/cmd"
	"net/http"
)

type VersionCtl struct {
}

func NewVersionCtl() *VersionCtl {
	return &VersionCtl{}
}

func (this *VersionCtl) Version(c *gin.Context) {
	rst := v1.NewResolve(cmd.NewVersionInfo())
	c.JSON(http.StatusOK, rst)
}

func (this *VersionCtl) Build(core *server.KCore) {
	core.Handle("GET", "/version", this.Version)
}
