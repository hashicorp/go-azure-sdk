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

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	capabilitiesClient, err := capabilities.NewCapabilitiesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Capabilities client: %+v", err)
	}
	configureFunc(capabilitiesClient.Client)

	capabilityTypesClient, err := capabilitytypes.NewCapabilityTypesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CapabilityTypes client: %+v", err)
	}
	configureFunc(capabilityTypesClient.Client)

	experimentsClient, err := experiments.NewExperimentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Experiments client: %+v", err)
	}
	configureFunc(experimentsClient.Client)

	targetTypesClient, err := targettypes.NewTargetTypesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TargetTypes client: %+v", err)
	}
	configureFunc(targetTypesClient.Client)

	targetsClient, err := targets.NewTargetsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Targets client: %+v", err)
	}
	configureFunc(targetsClient.Client)

	return &Client{
		Capabilities:    capabilitiesClient,
		CapabilityTypes: capabilityTypesClient,
		Experiments:     experimentsClient,
		TargetTypes:     targetTypesClient,
		Targets:         targetsClient,
	}, nil
}
