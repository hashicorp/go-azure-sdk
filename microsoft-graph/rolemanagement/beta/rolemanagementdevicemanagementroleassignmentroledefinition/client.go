package rolemanagementdevicemanagementroleassignmentroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDeviceManagementRoleAssignmentRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDeviceManagementRoleAssignmentRoleDefinitionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDeviceManagementRoleAssignmentRoleDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdevicemanagementroleassignmentroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDeviceManagementRoleAssignmentRoleDefinitionClient: %+v", err)
	}

	return &RoleManagementDeviceManagementRoleAssignmentRoleDefinitionClient{
		Client: client,
	}, nil
}
