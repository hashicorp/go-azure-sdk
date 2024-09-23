package devicecustomattributeshellscriptuserrunstatedevicerunstatemanageddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDeviceClient struct {
	Client *msgraph.Client
}

func NewDeviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecustomattributeshellscriptuserrunstatedevicerunstatemanageddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDeviceClient: %+v", err)
	}

	return &DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDeviceClient{
		Client: client,
	}, nil
}
