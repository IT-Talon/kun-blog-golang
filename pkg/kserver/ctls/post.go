package ctls

import (
	"github.com/gin-gonic/gin"
	"kun-blog-golang/core/server"
	"kun-blog-golang/pkg/store/fsstore"
	. "kun-blog-golang/pkg/utils"
	_ "kun-blog-golang/pkg/utils"
	"log"
)

type PostCtl struct {
	fs          *fsstore.FSStore
	articlePath string
}

func NewPostCtl() *PostCtl {
	articlePath := "./public/contents/articles"
	return &PostCtl{
		fs:          fsstore.NewFSStore(articlePath),
		articlePath: articlePath,
	}
}

func (this *PostCtl) Apply(c *gin.Context) {
	form, err := c.MultipartForm()
	CheckError(err)
	log.Println("apply", form)
	for inputName, fileList := range form.File {
		log.Println("处理文件：", inputName)
		f, err := fileList[0].Open()
		CheckError(err)
		defer f.Close()
		postM, err := ParseFileToPost(f)
		CheckError(err)
		postM.FileName = fileList[0].Filename
		CheckError(this.fs.Save(postM))
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "操作成功",
		"data": "",
	})

}

func (this *PostCtl) Build(core *server.KCore) {
	core.Handle("POST", "/post/apply", this.Apply)
}
