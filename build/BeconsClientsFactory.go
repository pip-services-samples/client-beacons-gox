package build

import (
	bclients "github.com/pip-services-samples/client-beacons-gox/clients/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type BeaconsClientFactory struct {
	cbuild.Factory
	NullClientDescriptor   *cref.Descriptor
	DirectClientDescriptor *cref.Descriptor
	HttpClientDescriptor   *cref.Descriptor
	MemoryClientDescriptor *cref.Descriptor
}

func NewBeaconsClientFactory() *BeaconsClientFactory {

	bcf := BeaconsClientFactory{}
	bcf.Factory = *cbuild.NewFactory()

	bcf.NullClientDescriptor = cref.NewDescriptor("beacons", "client", "null", "*", "1.0")
	bcf.DirectClientDescriptor = cref.NewDescriptor("beacons", "client", "direct", "*", "1.0")
	bcf.HttpClientDescriptor = cref.NewDescriptor("beacons", "client", "http", "*", "1.0")
	bcf.MemoryClientDescriptor = cref.NewDescriptor("beacons", "client", "memory", "*", "1.0")

	bcf.RegisterType(bcf.NullClientDescriptor, bclients.NewBeaconsNullClientV1)
	bcf.RegisterType(bcf.DirectClientDescriptor, bclients.NewBeaconsDirectClientV1)
	bcf.RegisterType(bcf.HttpClientDescriptor, bclients.NewBeaconsHttpClientV1)
	bcf.RegisterType(bcf.MemoryClientDescriptor, bclients.NewBeaconsMemoryClientV1)

	return &bcf
}
