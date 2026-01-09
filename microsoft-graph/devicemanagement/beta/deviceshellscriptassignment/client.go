package deviceshellscriptassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceShellScriptAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceShellScriptAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceShellScriptAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceshellscriptassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceShellScriptAssignmentClient: %+v", err)
	}

	return &DeviceShellScriptAssignmentClient{
		Client: client,
	}, nil
}
