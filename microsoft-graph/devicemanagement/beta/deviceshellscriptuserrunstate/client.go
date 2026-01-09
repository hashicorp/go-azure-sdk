package deviceshellscriptuserrunstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceShellScriptUserRunStateClient struct {
	Client *msgraph.Client
}

func NewDeviceShellScriptUserRunStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceShellScriptUserRunStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceshellscriptuserrunstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceShellScriptUserRunStateClient: %+v", err)
	}

	return &DeviceShellScriptUserRunStateClient{
		Client: client,
	}, nil
}
