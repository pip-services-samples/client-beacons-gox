package test_clients1

import (
	"context"
	"testing"

	clients1 "github.com/pip-services-samples/client-beacons-gox/clients/version1"
	logic "github.com/pip-services-samples/service-beacons-gox/logic"
	persist "github.com/pip-services-samples/service-beacons-gox/persistence"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
)

type beaconsDirectClientV1Test struct {
	persistence *persist.BeaconsMemoryPersistence
	controller  *logic.BeaconsController
	client      *clients1.BeaconsDirectClientV1
	fixture     *BeaconsClientV1Fixture
	ctx         context.Context
}

func newBeaconsDirectClientV1Test() *beaconsDirectClientV1Test {
	ctx := context.Background()
	persistence := persist.NewBeaconsMemoryPersistence()
	persistence.Configure(ctx, cconf.NewEmptyConfigParams())

	controller := logic.NewBeaconsController()
	controller.Configure(ctx, cconf.NewEmptyConfigParams())

	client := clients1.NewBeaconsDirectClientV1()
	client.Configure(ctx, cconf.NewEmptyConfigParams())

	references := cref.NewReferencesFromTuples(ctx,
		cref.NewDescriptor("beacons", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("beacons", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("beacons", "client", "direct", "default", "1.0"), client,
	)
	controller.SetReferences(ctx, references)
	client.SetReferences(ctx, references)

	fixture := NewBeaconsClientV1Fixture(client)

	return &beaconsDirectClientV1Test{
		persistence: persistence,
		controller:  controller,
		client:      client,
		fixture:     fixture,
		ctx:         ctx,
	}
}

func (c *beaconsDirectClientV1Test) setup(t *testing.T) {
	err := c.persistence.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open persistence", err)
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

func (c *beaconsDirectClientV1Test) teardown(t *testing.T) {
	err := c.client.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close client", err)
	}

	err = c.persistence.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestBeaconsDirectClientV1(t *testing.T) {
	c := newBeaconsDirectClientV1Test()

	c.setup(t)
	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)
	t.Run("Calculate Positions", c.fixture.TestCalculatePosition)
	c.teardown(t)
}
