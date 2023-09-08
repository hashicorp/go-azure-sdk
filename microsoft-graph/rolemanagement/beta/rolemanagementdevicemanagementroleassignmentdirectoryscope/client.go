package rolemanagementdevicemanagementroleassignmentdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDeviceManagementRoleAssignmentDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDeviceManagementRoleAssignmentDirectoryScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDeviceManagementRoleAssignmentDirectoryScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdevicemanagementroleassignmentdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDeviceManagementRoleAssignmentDirectoryScopeClient: %+v", err)
	}

	return &RoleManagementDeviceManagementRoleAssignmentDirectoryScopeClient{
		Client: client,
	}, nil
}
