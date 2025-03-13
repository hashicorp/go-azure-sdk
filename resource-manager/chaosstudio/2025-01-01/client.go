package v2025_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2025-01-01/capabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2025-01-01/capabilitytypes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2025-01-01/experimentexecutions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2025-01-01/experiments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2025-01-01/targets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2025-01-01/targettypes"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Capabilities         *capabilities.CapabilitiesClient
	CapabilityTypes      *capabilitytypes.CapabilityTypesClient
	ExperimentExecutions *experimentexecutions.ExperimentExecutionsClient
	Experiments          *experiments.ExperimentsClient
	TargetTypes          *targettypes.TargetTypesClient
	Targets              *targets.TargetsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
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

	experimentExecutionsClient, err := experimentexecutions.NewExperimentExecutionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExperimentExecutions client: %+v", err)
	}
	configureFunc(experimentExecutionsClient.Client)

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
		Capabilities:         capabilitiesClient,
		CapabilityTypes:      capabilityTypesClient,
		ExperimentExecutions: experimentExecutionsClient,
		Experiments:          experimentsClient,
		TargetTypes:          targetTypesClient,
		Targets:              targetsClient,
	}, nil
}
