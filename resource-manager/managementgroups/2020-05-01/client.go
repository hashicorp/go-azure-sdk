package v2020_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2020-05-01/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2020-05-01/entities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2020-05-01/managementgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2020-05-01/tenantbackfill"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CheckNameAvailability *checknameavailability.CheckNameAvailabilityClient
	Entities              *entities.EntitiesClient
	ManagementGroups      *managementgroups.ManagementGroupsClient
	TenantBackfill        *tenantbackfill.TenantBackfillClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	checkNameAvailabilityClient, err := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailability client: %+v", err)
	}
	configureFunc(checkNameAvailabilityClient.Client)

	entitiesClient, err := entities.NewEntitiesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Entities client: %+v", err)
	}
	configureFunc(entitiesClient.Client)

	managementGroupsClient, err := managementgroups.NewManagementGroupsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ManagementGroups client: %+v", err)
	}
	configureFunc(managementGroupsClient.Client)

	tenantBackfillClient, err := tenantbackfill.NewTenantBackfillClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TenantBackfill client: %+v", err)
	}
	configureFunc(tenantBackfillClient.Client)

	return &Client{
		CheckNameAvailability: checkNameAvailabilityClient,
		Entities:              entitiesClient,
		ManagementGroups:      managementGroupsClient,
		TenantBackfill:        tenantBackfillClient,
	}, nil
}
