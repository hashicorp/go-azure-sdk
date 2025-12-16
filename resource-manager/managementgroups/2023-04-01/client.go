package v2023_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2023-04-01/hierarchysettingsoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2023-04-01/managementgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2023-04-01/managements"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2023-04-01/subscriptionundermanagementgroups"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	HierarchySettingsOperationGroup   *hierarchysettingsoperationgroup.HierarchySettingsOperationGroupClient
	ManagementGroups                  *managementgroups.ManagementGroupsClient
	Managements                       *managements.ManagementsClient
	SubscriptionUnderManagementGroups *subscriptionundermanagementgroups.SubscriptionUnderManagementGroupsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	hierarchySettingsOperationGroupClient, err := hierarchysettingsoperationgroup.NewHierarchySettingsOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HierarchySettingsOperationGroup client: %+v", err)
	}
	configureFunc(hierarchySettingsOperationGroupClient.Client)

	managementGroupsClient, err := managementgroups.NewManagementGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagementGroups client: %+v", err)
	}
	configureFunc(managementGroupsClient.Client)

	managementsClient, err := managements.NewManagementsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Managements client: %+v", err)
	}
	configureFunc(managementsClient.Client)

	subscriptionUnderManagementGroupsClient, err := subscriptionundermanagementgroups.NewSubscriptionUnderManagementGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubscriptionUnderManagementGroups client: %+v", err)
	}
	configureFunc(subscriptionUnderManagementGroupsClient.Client)

	return &Client{
		HierarchySettingsOperationGroup:   hierarchySettingsOperationGroupClient,
		ManagementGroups:                  managementGroupsClient,
		Managements:                       managementsClient,
		SubscriptionUnderManagementGroups: subscriptionUnderManagementGroupsClient,
	}, nil
}
