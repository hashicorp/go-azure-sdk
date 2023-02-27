// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_10_01

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
