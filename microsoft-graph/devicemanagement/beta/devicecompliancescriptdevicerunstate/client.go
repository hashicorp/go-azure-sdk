package devicecompliancescriptdevicerunstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptDeviceRunStateClient struct {
	Client *msgraph.Client
}

func NewDeviceComplianceScriptDeviceRunStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceComplianceScriptDeviceRunStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancescriptdevicerunstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceComplianceScriptDeviceRunStateClient: %+v", err)
	}

	return &DeviceComplianceScriptDeviceRunStateClient{
		Client: client,
	}, nil
}
