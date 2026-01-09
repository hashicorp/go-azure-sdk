package devicemanagementroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementRoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementRoleAssignmentClient: %+v", err)
	}

	return &DeviceManagementRoleAssignmentClient{
		Client: client,
	}, nil
}
