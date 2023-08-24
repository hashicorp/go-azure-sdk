package v2015_11_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/operationsmanagement/2015-11-01-preview/managementassociation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationsmanagement/2015-11-01-preview/managementconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationsmanagement/2015-11-01-preview/solution"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ManagementAssociation   *managementassociation.ManagementAssociationClient
	ManagementConfiguration *managementconfiguration.ManagementConfigurationClient
	Solution                *solution.SolutionClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	managementAssociationClient, err := managementassociation.NewManagementAssociationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagementAssociation client: %+v", err)
	}
	configureFunc(managementAssociationClient.Client)

	managementConfigurationClient, err := managementconfiguration.NewManagementConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagementConfiguration client: %+v", err)
	}
	configureFunc(managementConfigurationClient.Client)

	solutionClient, err := solution.NewSolutionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Solution client: %+v", err)
	}
	configureFunc(solutionClient.Client)

	return &Client{
		ManagementAssociation:   managementAssociationClient,
		ManagementConfiguration: managementConfigurationClient,
		Solution:                solutionClient,
	}, nil
}
