package devicecompliancescriptdevicerunstatemanageddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptDeviceRunStateManagedDeviceClient struct {
	Client *msgraph.Client
}

func NewDeviceComplianceScriptDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceComplianceScriptDeviceRunStateManagedDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancescriptdevicerunstatemanageddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceComplianceScriptDeviceRunStateManagedDeviceClient: %+v", err)
	}

	return &DeviceComplianceScriptDeviceRunStateManagedDeviceClient{
		Client: client,
	}, nil
}
