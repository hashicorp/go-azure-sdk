package v2021_01_01

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/powerbidedicated/2021-01-01/autoscalevcores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/powerbidedicated/2021-01-01/capacities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/powerbidedicated/2021-01-01/powerbidedicated"
)

type Client struct {
	AutoScaleVCores  *autoscalevcores.AutoScaleVCoresClient
	Capacities       *capacities.CapacitiesClient
	PowerBIDedicated *powerbidedicated.PowerBIDedicatedClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	autoScaleVCoresClient := autoscalevcores.NewAutoScaleVCoresClientWithBaseURI(endpoint)
	configureAuthFunc(&autoScaleVCoresClient.Client)

	capacitiesClient := capacities.NewCapacitiesClientWithBaseURI(endpoint)
	configureAuthFunc(&capacitiesClient.Client)

	powerBIDedicatedClient := powerbidedicated.NewPowerBIDedicatedClientWithBaseURI(endpoint)
	configureAuthFunc(&powerBIDedicatedClient.Client)

	return Client{
		AutoScaleVCores:  &autoScaleVCoresClient,
		Capacities:       &capacitiesClient,
		PowerBIDedicated: &powerBIDedicatedClient,
	}
}
