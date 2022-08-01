package v1

import "kun-blog-golang/pkg/kctl/clientset/rest"

type V1Client struct {
	client *rest.RESTClient
}

func New(client *rest.RESTClient) *V1Client {
	return &V1Client{
		client: client,
	}
}

var (
	_ V1Interface   = &V1Client{}
	_ VersionGetter = &V1Client{}
)

func (this *V1Client) Version() VersionInterface {
	return newVersion(this.client)
}
