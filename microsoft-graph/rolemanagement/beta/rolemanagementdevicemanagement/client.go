package rolemanagementdevicemanagement

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDeviceManagementClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDeviceManagementClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDeviceManagementClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdevicemanagement", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDeviceManagementClient: %+v", err)
	}

	return &RoleManagementDeviceManagementClient{
		Client: client,
	}, nil
}
