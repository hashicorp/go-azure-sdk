package v2025_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2025-03-01/standbycontainergrouppoolruntimeviews"
	"github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2025-03-01/standbycontainergrouppools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2025-03-01/standbyvirtualmachinepoolruntimeviews"
	"github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2025-03-01/standbyvirtualmachinepools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2025-03-01/standbyvirtualmachines"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	StandbyContainerGroupPoolRuntimeViews *standbycontainergrouppoolruntimeviews.StandbyContainerGroupPoolRuntimeViewsClient
	StandbyContainerGroupPools            *standbycontainergrouppools.StandbyContainerGroupPoolsClient
	StandbyVirtualMachinePoolRuntimeViews *standbyvirtualmachinepoolruntimeviews.StandbyVirtualMachinePoolRuntimeViewsClient
	StandbyVirtualMachinePools            *standbyvirtualmachinepools.StandbyVirtualMachinePoolsClient
	StandbyVirtualMachines                *standbyvirtualmachines.StandbyVirtualMachinesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	standbyContainerGroupPoolRuntimeViewsClient, err := standbycontainergrouppoolruntimeviews.NewStandbyContainerGroupPoolRuntimeViewsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StandbyContainerGroupPoolRuntimeViews client: %+v", err)
	}
	configureFunc(standbyContainerGroupPoolRuntimeViewsClient.Client)

	standbyContainerGroupPoolsClient, err := standbycontainergrouppools.NewStandbyContainerGroupPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StandbyContainerGroupPools client: %+v", err)
	}
	configureFunc(standbyContainerGroupPoolsClient.Client)

	standbyVirtualMachinePoolRuntimeViewsClient, err := standbyvirtualmachinepoolruntimeviews.NewStandbyVirtualMachinePoolRuntimeViewsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StandbyVirtualMachinePoolRuntimeViews client: %+v", err)
	}
	configureFunc(standbyVirtualMachinePoolRuntimeViewsClient.Client)

	standbyVirtualMachinePoolsClient, err := standbyvirtualmachinepools.NewStandbyVirtualMachinePoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StandbyVirtualMachinePools client: %+v", err)
	}
	configureFunc(standbyVirtualMachinePoolsClient.Client)

	standbyVirtualMachinesClient, err := standbyvirtualmachines.NewStandbyVirtualMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StandbyVirtualMachines client: %+v", err)
	}
	configureFunc(standbyVirtualMachinesClient.Client)

	return &Client{
		StandbyContainerGroupPoolRuntimeViews: standbyContainerGroupPoolRuntimeViewsClient,
		StandbyContainerGroupPools:            standbyContainerGroupPoolsClient,
		StandbyVirtualMachinePoolRuntimeViews: standbyVirtualMachinePoolRuntimeViewsClient,
		StandbyVirtualMachinePools:            standbyVirtualMachinePoolsClient,
		StandbyVirtualMachines:                standbyVirtualMachinesClient,
	}, nil
}
