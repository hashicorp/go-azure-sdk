package v2022_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2022-10-01/marketplaceregistrationdefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2022-10-01/registrationassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2022-10-01/registrationdefinitions"
)

type Client struct {
	MarketplaceRegistrationDefinitions *marketplaceregistrationdefinitions.MarketplaceRegistrationDefinitionsClient
	RegistrationAssignments            *registrationassignments.RegistrationAssignmentsClient
	RegistrationDefinitions            *registrationdefinitions.RegistrationDefinitionsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	marketplaceRegistrationDefinitionsClient := marketplaceregistrationdefinitions.NewMarketplaceRegistrationDefinitionsClientWithBaseURI(endpoint)
	configureAuthFunc(&marketplaceRegistrationDefinitionsClient.Client)

	registrationAssignmentsClient := registrationassignments.NewRegistrationAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&registrationAssignmentsClient.Client)

	registrationDefinitionsClient := registrationdefinitions.NewRegistrationDefinitionsClientWithBaseURI(endpoint)
	configureAuthFunc(&registrationDefinitionsClient.Client)

	return Client{
		MarketplaceRegistrationDefinitions: &marketplaceRegistrationDefinitionsClient,
		RegistrationAssignments:            &registrationAssignmentsClient,
		RegistrationDefinitions:            &registrationDefinitionsClient,
	}
}
