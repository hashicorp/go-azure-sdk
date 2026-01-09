package devicehealthscript

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptClient struct {
	Client *msgraph.Client
}

func NewDeviceHealthScriptClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceHealthScriptClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicehealthscript", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceHealthScriptClient: %+v", err)
	}

	return &DeviceHealthScriptClient{
		Client: client,
	}, nil
}
