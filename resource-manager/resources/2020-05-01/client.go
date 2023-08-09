package v2020_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-05-01/managementlocks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-05-01/privatelinkassociation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-05-01/resourcemanagementprivatelink"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ManagementLocks               *managementlocks.ManagementLocksClient
	PrivateLinkAssociation        *privatelinkassociation.PrivateLinkAssociationClient
	ResourceManagementPrivateLink *resourcemanagementprivatelink.ResourceManagementPrivateLinkClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	managementLocksClient, err := managementlocks.NewManagementLocksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagementLocks client: %+v", err)
	}
	configureFunc(managementLocksClient.Client)

	privateLinkAssociationClient, err := privatelinkassociation.NewPrivateLinkAssociationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkAssociation client: %+v", err)
	}
	configureFunc(privateLinkAssociationClient.Client)

	resourceManagementPrivateLinkClient, err := resourcemanagementprivatelink.NewResourceManagementPrivateLinkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceManagementPrivateLink client: %+v", err)
	}
	configureFunc(resourceManagementPrivateLinkClient.Client)

	return &Client{
		ManagementLocks:               managementLocksClient,
		PrivateLinkAssociation:        privateLinkAssociationClient,
		ResourceManagementPrivateLink: resourceManagementPrivateLinkClient,
	}, nil
}
