package v1

import (
	v1 "kun-blog-golang/pkg/apis/v1"
	"kun-blog-golang/pkg/kctl/clientset/rest"
)

type VersionInterface interface {
	Get() (rVersion *v1.Version, err error)
}

type version struct {
	client *rest.RESTClient
}

func (this *version) Get() (rVersion *v1.Version, err error) {
	rVersion = &v1.Version{}
	err = this.client.Get().Path("/v1/version").Do().Into(rVersion)
	return
}

func newVersion(client *rest.RESTClient) *version {
	return &version{
		client: client,
	}
}

var _ VersionInterface = &version{}
