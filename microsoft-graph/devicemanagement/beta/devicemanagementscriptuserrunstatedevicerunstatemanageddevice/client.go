package devicemanagementscriptuserrunstatedevicerunstatemanageddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptUserRunStateDeviceRunStateManagedDeviceClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementScriptUserRunStateDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementScriptUserRunStateDeviceRunStateManagedDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementscriptuserrunstatedevicerunstatemanageddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementScriptUserRunStateDeviceRunStateManagedDeviceClient: %+v", err)
	}

	return &DeviceManagementScriptUserRunStateDeviceRunStateManagedDeviceClient{
		Client: client,
	}, nil
}
