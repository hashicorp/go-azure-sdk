package devicemanagementscriptdevicerunstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptDeviceRunStateClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementScriptDeviceRunStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementScriptDeviceRunStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementscriptdevicerunstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementScriptDeviceRunStateClient: %+v", err)
	}

	return &DeviceManagementScriptDeviceRunStateClient{
		Client: client,
	}, nil
}
