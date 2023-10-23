package v2022_05_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2022-05-01-preview/permissions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2022-05-01-preview/roledefinitions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Permissions     *permissions.PermissionsClient
	RoleDefinitions *roledefinitions.RoleDefinitionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	permissionsClient, err := permissions.NewPermissionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Permissions client: %+v", err)
	}
	configureFunc(permissionsClient.Client)

	roleDefinitionsClient, err := roledefinitions.NewRoleDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleDefinitions client: %+v", err)
	}
	configureFunc(roleDefinitionsClient.Client)

	return &Client{
		Permissions:     permissionsClient,
		RoleDefinitions: roleDefinitionsClient,
	}, nil
}
