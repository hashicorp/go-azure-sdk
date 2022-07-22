package v2022_07_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/capabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/capabilitytypes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/experiments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/targets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/targettypes"
)

type Client struct {
	Capabilities    *capabilities.CapabilitiesClient
	CapabilityTypes *capabilitytypes.CapabilityTypesClient
	Experiments     *experiments.ExperimentsClient
	TargetTypes     *targettypes.TargetTypesClient
	Targets         *targets.TargetsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	capabilitiesClient := capabilities.NewCapabilitiesClientWithBaseURI(endpoint)
	configureAuthFunc(&capabilitiesClient.Client)

	capabilityTypesClient := capabilitytypes.NewCapabilityTypesClientWithBaseURI(endpoint)
	configureAuthFunc(&capabilityTypesClient.Client)

	experimentsClient := experiments.NewExperimentsClientWithBaseURI(endpoint)
	configureAuthFunc(&experimentsClient.Client)

	targetTypesClient := targettypes.NewTargetTypesClientWithBaseURI(endpoint)
	configureAuthFunc(&targetTypesClient.Client)

	targetsClient := targets.NewTargetsClientWithBaseURI(endpoint)
	configureAuthFunc(&targetsClient.Client)

	return Client{
		Capabilities:    &capabilitiesClient,
		CapabilityTypes: &capabilityTypesClient,
		Experiments:     &experimentsClient,
		TargetTypes:     &targetTypesClient,
		Targets:         &targetsClient,
	}
}
