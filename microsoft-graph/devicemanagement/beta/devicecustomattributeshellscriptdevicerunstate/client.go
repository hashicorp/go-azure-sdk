package devicecustomattributeshellscriptdevicerunstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeShellScriptDeviceRunStateClient struct {
	Client *msgraph.Client
}

func NewDeviceCustomAttributeShellScriptDeviceRunStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCustomAttributeShellScriptDeviceRunStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecustomattributeshellscriptdevicerunstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCustomAttributeShellScriptDeviceRunStateClient: %+v", err)
	}

	return &DeviceCustomAttributeShellScriptDeviceRunStateClient{
		Client: client,
	}, nil
}
