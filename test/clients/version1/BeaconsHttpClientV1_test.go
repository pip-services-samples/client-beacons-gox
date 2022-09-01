package test_clients1

import (
	"context"
	"testing"

	clients1 "github.com/pip-services-samples/client-beacons-gox/clients/version1"
	logic "github.com/pip-services-samples/service-beacons-gox/logic"
	persist "github.com/pip-services-samples/service-beacons-gox/persistence"
	services1 "github.com/pip-services-samples/service-beacons-gox/services/version1"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
)

type beaconsCommandableHttpClientV1Test struct {
	persistence *persist.BeaconsMemoryPersistence
	controller  *logic.BeaconsController
	service     *services1.BeaconsHttpServiceV1
	client      *clients1.BeaconsHttpClientV1
	fixture     *BeaconsClientV1Fixture
	ctx         context.Context
}

func newBeaconsHttpClientV1Test() *beaconsCommandableHttpClientV1Test {
	ctx := context.Background()
	persistence := persist.NewBeaconsMemoryPersistence()
	persistence.Configure(ctx, cconf.NewEmptyConfigParams())

	controller := logic.NewBeaconsController()
	controller.Configure(ctx, cconf.NewEmptyConfigParams())

	httpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3000",
		"connection.host", "localhost",
	)

	service := services1.NewBeaconsHttpServiceV1()
	service.Configure(ctx, httpConfig)

	client := clients1.NewBeaconsHttpClientV1()
	client.Configure(ctx, httpConfig)

	references := cref.NewReferencesFromTuples(ctx,
		cref.NewDescriptor("beacons", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("beacons", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("beacons", "service", "http", "default", "1.0"), service,
		cref.NewDescriptor("beacons", "client", "http", "default", "1.0"), client,
	)
	controller.SetReferences(ctx, references)
	service.SetReferences(ctx, references)
	client.SetReferences(ctx, references)

	fixture := NewBeaconsClientV1Fixture(client)

	return &beaconsCommandableHttpClientV1Test{
		persistence: persistence,
		controller:  controller,
		service:     service,
		client:      client,
		fixture:     fixture,
		ctx:         ctx,
	}
}

func (c *beaconsCommandableHttpClientV1Test) setup(t *testing.T) {
	err := c.persistence.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.service.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open service", err)
	}

	err = c.client.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open client", err)
	}

	err = c.persistence.Clear(c.ctx, "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *beaconsCommandableHttpClientV1Test) teardown(t *testing.T) {
	err := c.client.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close client", err)
	}

	err = c.service.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close service", err)
	}

	err = c.persistence.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestBeaconsHttpClientV1(t *testing.T) {
	c := newBeaconsHttpClientV1Test()

	c.setup(t)

	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)

	t.Run("Calculate Positions", c.fixture.TestCalculatePosition)
	c.teardown(t)
}
