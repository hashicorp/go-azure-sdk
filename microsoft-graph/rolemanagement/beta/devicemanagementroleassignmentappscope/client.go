package devicemanagementroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementRoleAssignmentAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementRoleAssignmentAppScopeClient: %+v", err)
	}

	return &DeviceManagementRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
