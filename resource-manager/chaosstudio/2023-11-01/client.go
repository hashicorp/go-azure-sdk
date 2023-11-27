package v2023_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2023-11-01/asyncoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2023-11-01/capabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2023-11-01/capabilitytypes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2023-11-01/experiments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2023-11-01/targets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2023-11-01/targettypes"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AsyncOperations *asyncoperations.AsyncOperationsClient
	Capabilities    *capabilities.CapabilitiesClient
	CapabilityTypes *capabilitytypes.CapabilityTypesClient
	Experiments     *experiments.ExperimentsClient
	TargetTypes     *targettypes.TargetTypesClient
	Targets         *targets.TargetsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	asyncOperationsClient, err := asyncoperations.NewAsyncOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AsyncOperations client: %+v", err)
	}
	configureFunc(asyncOperationsClient.Client)

	capabilitiesClient, err := capabilities.NewCapabilitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Capabilities client: %+v", err)
	}
	configureFunc(capabilitiesClient.Client)

	capabilityTypesClient, err := capabilitytypes.NewCapabilityTypesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapabilityTypes client: %+v", err)
	}
	configureFunc(capabilityTypesClient.Client)

	experimentsClient, err := experiments.NewExperimentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Experiments client: %+v", err)
	}
	configureFunc(experimentsClient.Client)

	targetTypesClient, err := targettypes.NewTargetTypesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TargetTypes client: %+v", err)
	}
	configureFunc(targetTypesClient.Client)

	targetsClient, err := targets.NewTargetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Targets client: %+v", err)
	}
	configureFunc(targetsClient.Client)

	return &Client{
		AsyncOperations: asyncOperationsClient,
		Capabilities:    capabilitiesClient,
		CapabilityTypes: capabilityTypesClient,
		Experiments:     experimentsClient,
		TargetTypes:     targetTypesClient,
		Targets:         targetsClient,
	}, nil
}
