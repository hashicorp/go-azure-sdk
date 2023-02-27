package v2019_06_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2019-06-01/registrationassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2019-06-01/registrationdefinitions"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
