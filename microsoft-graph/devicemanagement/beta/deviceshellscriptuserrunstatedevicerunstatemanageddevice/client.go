package deviceshellscriptuserrunstatedevicerunstatemanageddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceShellScriptUserRunStateDeviceRunStateManagedDeviceClient struct {
	Client *msgraph.Client
}

func NewDeviceShellScriptUserRunStateDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceShellScriptUserRunStateDeviceRunStateManagedDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceshellscriptuserrunstatedevicerunstatemanageddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceShellScriptUserRunStateDeviceRunStateManagedDeviceClient: %+v", err)
	}

	return &DeviceShellScriptUserRunStateDeviceRunStateManagedDeviceClient{
		Client: client,
	}, nil
}
