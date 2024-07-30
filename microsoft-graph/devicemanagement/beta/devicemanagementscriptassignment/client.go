package devicemanagementscriptassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementScriptAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementScriptAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementscriptassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementScriptAssignmentClient: %+v", err)
	}

	return &DeviceManagementScriptAssignmentClient{
		Client: client,
	}, nil
}
