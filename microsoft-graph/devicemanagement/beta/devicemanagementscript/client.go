package devicemanagementscript

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementScriptClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementScriptClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementscript", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementScriptClient: %+v", err)
	}

	return &DeviceManagementScriptClient{
		Client: client,
	}, nil
}
