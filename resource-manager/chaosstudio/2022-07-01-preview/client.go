package v2022_07_01_preview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/capabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/capabilitytypes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/experiments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/targets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/targettypes"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Capabilities    *capabilities.CapabilitiesClient
	CapabilityTypes *capabilitytypes.CapabilityTypesClient
	Experiments     *experiments.ExperimentsClient
	TargetTypes     *targettypes.TargetTypesClient
	Targets         *targets.TargetsClient
}

func NewClientWithBaseURI(api environments.Api, configureAuthFunc func(c *resourcemanager.Client)) (*Client, error) {
	capabilitiesClient, err := capabilities.NewCapabilitiesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Capabilities: %+v", err)
	}
	configureAuthFunc(capabilitiesClient.Client)

	capabilityTypesClient, err := capabilitytypes.NewCapabilityTypesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for CapabilityTypes: %+v", err)
	}
	configureAuthFunc(capabilityTypesClient.Client)

	experimentsClient, err := experiments.NewExperimentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Experiments: %+v", err)
	}
	configureAuthFunc(experimentsClient.Client)

	targetTypesClient, err := targettypes.NewTargetTypesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for TargetTypes: %+v", err)
	}
	configureAuthFunc(targetTypesClient.Client)

	targetsClient, err := targets.NewTargetsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Targets: %+v", err)
	}
	configureAuthFunc(targetsClient.Client)

	return &Client{
		Capabilities:    capabilitiesClient,
		CapabilityTypes: capabilityTypesClient,
		Experiments:     experimentsClient,
		TargetTypes:     targetTypesClient,
		Targets:         targetsClient,
	}, nil
}
