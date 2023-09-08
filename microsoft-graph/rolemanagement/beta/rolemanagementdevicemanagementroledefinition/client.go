package rolemanagementdevicemanagementroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDeviceManagementRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDeviceManagementRoleDefinitionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDeviceManagementRoleDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdevicemanagementroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDeviceManagementRoleDefinitionClient: %+v", err)
	}

	return &RoleManagementDeviceManagementRoleDefinitionClient{
		Client: client,
	}, nil
}
