package clients1

import (
	"context"

	data1 "github.com/pip-services-samples/service-beacons-gox/data/version1"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type BeaconsNullClientV1 struct {
}

func NewBeaconsNullClientV1() *BeaconsNullClientV1 {
	return &BeaconsNullClientV1{}
}

func (c *BeaconsNullClientV1) GetBeacons(ctx context.Context,
	correlationId string, filter cdata.FilterParams,
	paging cdata.PagingParams) (*cdata.DataPage[data1.BeaconV1], error) {
	return cdata.NewEmptyDataPage[data1.BeaconV1](), nil
}

func (c *BeaconsNullClientV1) GetBeaconById(ctx context.Context,
	correlationId string, beaconId string) (*data1.BeaconV1, error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) GetBeaconByUdi(ctx context.Context,
	correlationId string, udi string) (*data1.BeaconV1, error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) CalculatePosition(ctx context.Context,
	correlationId string, siteId string, udis []string) (*data1.GeoPointV1, error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) CreateBeacon(ctx context.Context,
	correlationId string, beacon *data1.BeaconV1) (*data1.BeaconV1, error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) UpdateBeacon(ctx context.Context,
	correlationId string, beacon *data1.BeaconV1) (*data1.BeaconV1, error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) DeleteBeaconById(ctx context.Context,
	correlationId string, beaconId string) (*data1.BeaconV1, error) {
	return nil, nil
}
