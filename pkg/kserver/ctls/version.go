package ctls

import (
	"github.com/gin-gonic/gin"
	"kun-blog-golang/core/server"
	"kun-blog-golang/pkg/kctl/cmd"
	"net/http"
)

type VersionCtl struct {
}

func NewVersionCtl() *VersionCtl {
	return &VersionCtl{}
}

func (this *VersionCtl) Version(c *gin.Context) {
	version := cmd.NewVersionInfo()
	c.JSONP(http.StatusOK, gin.H{"Version": version.Version, "GoVersion": version.GoVersion})
}

func (this *VersionCtl) Build(core *server.KCore) {
	core.Handle("GET", "/version", this.Version)
}
