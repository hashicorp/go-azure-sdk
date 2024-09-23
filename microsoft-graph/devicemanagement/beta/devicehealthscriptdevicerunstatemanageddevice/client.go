package devicehealthscriptdevicerunstatemanageddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptDeviceRunStateManagedDeviceClient struct {
	Client *msgraph.Client
}

func NewDeviceHealthScriptDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceHealthScriptDeviceRunStateManagedDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicehealthscriptdevicerunstatemanageddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceHealthScriptDeviceRunStateManagedDeviceClient: %+v", err)
	}

	return &DeviceHealthScriptDeviceRunStateManagedDeviceClient{
		Client: client,
	}, nil
}
