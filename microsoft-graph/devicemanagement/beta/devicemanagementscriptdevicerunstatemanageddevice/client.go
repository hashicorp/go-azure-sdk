package devicemanagementscriptdevicerunstatemanageddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptDeviceRunStateManagedDeviceClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementScriptDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementScriptDeviceRunStateManagedDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementscriptdevicerunstatemanageddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementScriptDeviceRunStateManagedDeviceClient: %+v", err)
	}

	return &DeviceManagementScriptDeviceRunStateManagedDeviceClient{
		Client: client,
	}, nil
}
