// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2019_06_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2019-06-01/registrationassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2019-06-01/registrationdefinitions"
)

type Client struct {
	RegistrationAssignments *registrationassignments.RegistrationAssignmentsClient
	RegistrationDefinitions *registrationdefinitions.RegistrationDefinitionsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	registrationAssignmentsClient := registrationassignments.NewRegistrationAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&registrationAssignmentsClient.Client)

	registrationDefinitionsClient := registrationdefinitions.NewRegistrationDefinitionsClientWithBaseURI(endpoint)
	configureAuthFunc(&registrationDefinitionsClient.Client)

	return Client{
		RegistrationAssignments: &registrationAssignmentsClient,
		RegistrationDefinitions: &registrationDefinitionsClient,
	}
}
