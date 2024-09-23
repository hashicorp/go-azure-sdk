package devicemanagementroledefinitioninheritspermissionsfrom

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRoleDefinitionInheritsPermissionsFromClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementRoleDefinitionInheritsPermissionsFromClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementroledefinitioninheritspermissionsfrom", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementRoleDefinitionInheritsPermissionsFromClient: %+v", err)
	}

	return &DeviceManagementRoleDefinitionInheritsPermissionsFromClient{
		Client: client,
	}, nil
}
