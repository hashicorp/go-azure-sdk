package v2015_11_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationsmanagement/2015-11-01-preview/managementassociation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationsmanagement/2015-11-01-preview/managementconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationsmanagement/2015-11-01-preview/solution"
)

type Client struct {
	ManagementAssociation   *managementassociation.ManagementAssociationClient
	ManagementConfiguration *managementconfiguration.ManagementConfigurationClient
	Solution                *solution.SolutionClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	managementAssociationClient := managementassociation.NewManagementAssociationClientWithBaseURI(endpoint)
	configureAuthFunc(&managementAssociationClient.Client)

	managementConfigurationClient := managementconfiguration.NewManagementConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&managementConfigurationClient.Client)

	solutionClient := solution.NewSolutionClientWithBaseURI(endpoint)
	configureAuthFunc(&solutionClient.Client)

	return Client{
		ManagementAssociation:   &managementAssociationClient,
		ManagementConfiguration: &managementConfigurationClient,
		Solution:                &solutionClient,
	}
}
