package devicecustomattributeshellscriptdevicerunstatemanageddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeShellScriptDeviceRunStateManagedDeviceClient struct {
	Client *msgraph.Client
}

func NewDeviceCustomAttributeShellScriptDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCustomAttributeShellScriptDeviceRunStateManagedDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecustomattributeshellscriptdevicerunstatemanageddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCustomAttributeShellScriptDeviceRunStateManagedDeviceClient: %+v", err)
	}

	return &DeviceCustomAttributeShellScriptDeviceRunStateManagedDeviceClient{
		Client: client,
	}, nil
}
