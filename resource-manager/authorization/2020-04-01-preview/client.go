package v2020_04_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-04-01-preview/roleassignments"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	RoleAssignments *roleassignments.RoleAssignmentsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	roleAssignmentsClient, err := roleassignments.NewRoleAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleAssignments client: %+v", err)
	}
	configureFunc(roleAssignmentsClient.Client)

	return &Client{
		RoleAssignments: roleAssignmentsClient,
	}, nil
}
