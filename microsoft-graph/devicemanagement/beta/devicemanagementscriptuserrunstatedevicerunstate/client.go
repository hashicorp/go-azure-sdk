package devicemanagementscriptuserrunstatedevicerunstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptUserRunStateDeviceRunStateClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementScriptUserRunStateDeviceRunStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementScriptUserRunStateDeviceRunStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementscriptuserrunstatedevicerunstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementScriptUserRunStateDeviceRunStateClient: %+v", err)
	}

	return &DeviceManagementScriptUserRunStateDeviceRunStateClient{
		Client: client,
	}, nil
}
