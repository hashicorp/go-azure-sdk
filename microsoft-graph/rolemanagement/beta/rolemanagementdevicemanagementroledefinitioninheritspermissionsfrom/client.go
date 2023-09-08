package rolemanagementdevicemanagementroledefinitioninheritspermissionsfrom

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdevicemanagementroledefinitioninheritspermissionsfrom", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClient: %+v", err)
	}

	return &RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClient{
		Client: client,
	}, nil
}
