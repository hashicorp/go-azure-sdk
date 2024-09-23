package devicecustomattributeshellscriptuserrunstatedevicerunstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateClient struct {
	Client *msgraph.Client
}

func NewDeviceCustomAttributeShellScriptUserRunStateDeviceRunStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecustomattributeshellscriptuserrunstatedevicerunstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateClient: %+v", err)
	}

	return &DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateClient{
		Client: client,
	}, nil
}
