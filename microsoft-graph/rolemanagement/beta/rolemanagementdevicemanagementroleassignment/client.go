package rolemanagementdevicemanagementroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDeviceManagementRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDeviceManagementRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDeviceManagementRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdevicemanagementroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDeviceManagementRoleAssignmentClient: %+v", err)
	}

	return &RoleManagementDeviceManagementRoleAssignmentClient{
		Client: client,
	}, nil
}
