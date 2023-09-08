package rolemanagementdirectorytransitiveroleassignmentdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryTransitiveRoleAssignmentDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryTransitiveRoleAssignmentDirectoryScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectorytransitiveroleassignmentdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryTransitiveRoleAssignmentDirectoryScopeClient: %+v", err)
	}

	return &RoleManagementDirectoryTransitiveRoleAssignmentDirectoryScopeClient{
		Client: client,
	}, nil
}
