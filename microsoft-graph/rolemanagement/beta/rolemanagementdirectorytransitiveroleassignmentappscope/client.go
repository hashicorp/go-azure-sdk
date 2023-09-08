package rolemanagementdirectorytransitiveroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryTransitiveRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryTransitiveRoleAssignmentAppScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryTransitiveRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectorytransitiveroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryTransitiveRoleAssignmentAppScopeClient: %+v", err)
	}

	return &RoleManagementDirectoryTransitiveRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
