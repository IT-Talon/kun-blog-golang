package v1

type V1Interface interface {
	VersionGetter
}

type VersionGetter interface {
	Version() VersionInterface
}
