package v2019_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2019-06-01/registrationassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2019-06-01/registrationdefinitions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	RegistrationAssignments *registrationassignments.RegistrationAssignmentsClient
	RegistrationDefinitions *registrationdefinitions.RegistrationDefinitionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	registrationAssignmentsClient, err := registrationassignments.NewRegistrationAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RegistrationAssignments client: %+v", err)
	}
	configureFunc(registrationAssignmentsClient.Client)

	registrationDefinitionsClient, err := registrationdefinitions.NewRegistrationDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RegistrationDefinitions client: %+v", err)
	}
	configureFunc(registrationDefinitionsClient.Client)

	return &Client{
		RegistrationAssignments: registrationAssignmentsClient,
		RegistrationDefinitions: registrationDefinitionsClient,
	}, nil
}
