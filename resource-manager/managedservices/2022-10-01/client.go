package v2022_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2022-10-01/marketplaceregistrationdefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2022-10-01/operations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2022-10-01/registrationassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2022-10-01/registrationdefinitions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	MarketplaceRegistrationDefinitions *marketplaceregistrationdefinitions.MarketplaceRegistrationDefinitionsClient
	Operations                         *operations.OperationsClient
	RegistrationAssignments            *registrationassignments.RegistrationAssignmentsClient
	RegistrationDefinitions            *registrationdefinitions.RegistrationDefinitionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	marketplaceRegistrationDefinitionsClient, err := marketplaceregistrationdefinitions.NewMarketplaceRegistrationDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MarketplaceRegistrationDefinitions client: %+v", err)
	}
	configureFunc(marketplaceRegistrationDefinitionsClient.Client)

	operationsClient, err := operations.NewOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Operations client: %+v", err)
	}
	configureFunc(operationsClient.Client)

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
		MarketplaceRegistrationDefinitions: marketplaceRegistrationDefinitionsClient,
		Operations:                         operationsClient,
		RegistrationAssignments:            registrationAssignmentsClient,
		RegistrationDefinitions:            registrationDefinitionsClient,
	}, nil
}
