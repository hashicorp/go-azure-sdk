package v2018_01_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2018-01-01-preview/permissions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2018-01-01-preview/provideroperationsmetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2018-01-01-preview/roleassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2018-01-01-preview/roledefinitions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Permissions                *permissions.PermissionsClient
	ProviderOperationsMetadata *provideroperationsmetadata.ProviderOperationsMetadataClient
	RoleAssignments            *roleassignments.RoleAssignmentsClient
	RoleDefinitions            *roledefinitions.RoleDefinitionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	permissionsClient, err := permissions.NewPermissionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Permissions client: %+v", err)
	}
	configureFunc(permissionsClient.Client)

	providerOperationsMetadataClient, err := provideroperationsmetadata.NewProviderOperationsMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProviderOperationsMetadata client: %+v", err)
	}
	configureFunc(providerOperationsMetadataClient.Client)

	roleAssignmentsClient, err := roleassignments.NewRoleAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleAssignments client: %+v", err)
	}
	configureFunc(roleAssignmentsClient.Client)

	roleDefinitionsClient, err := roledefinitions.NewRoleDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleDefinitions client: %+v", err)
	}
	configureFunc(roleDefinitionsClient.Client)

	return &Client{
		Permissions:                permissionsClient,
		ProviderOperationsMetadata: providerOperationsMetadataClient,
		RoleAssignments:            roleAssignmentsClient,
		RoleDefinitions:            roleDefinitionsClient,
	}, nil
}
