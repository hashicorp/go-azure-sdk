package deviceshellscript

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceShellScriptClient struct {
	Client *msgraph.Client
}

func NewDeviceShellScriptClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceShellScriptClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceshellscript", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceShellScriptClient: %+v", err)
	}

	return &DeviceShellScriptClient{
		Client: client,
	}, nil
}
