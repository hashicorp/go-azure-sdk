package devicemanagementscriptuserrunstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptUserRunStateClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementScriptUserRunStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementScriptUserRunStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementscriptuserrunstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementScriptUserRunStateClient: %+v", err)
	}

	return &DeviceManagementScriptUserRunStateClient{
		Client: client,
	}, nil
}
