package devicehealthscriptdevicerunstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptDeviceRunStateClient struct {
	Client *msgraph.Client
}

func NewDeviceHealthScriptDeviceRunStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceHealthScriptDeviceRunStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicehealthscriptdevicerunstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceHealthScriptDeviceRunStateClient: %+v", err)
	}

	return &DeviceHealthScriptDeviceRunStateClient{
		Client: client,
	}, nil
}
