package deviceshellscriptdevicerunstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceShellScriptDeviceRunStateClient struct {
	Client *msgraph.Client
}

func NewDeviceShellScriptDeviceRunStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceShellScriptDeviceRunStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceshellscriptdevicerunstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceShellScriptDeviceRunStateClient: %+v", err)
	}

	return &DeviceShellScriptDeviceRunStateClient{
		Client: client,
	}, nil
}
