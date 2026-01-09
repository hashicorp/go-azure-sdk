package devicecustomattributeshellscriptassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeShellScriptAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceCustomAttributeShellScriptAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCustomAttributeShellScriptAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecustomattributeshellscriptassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCustomAttributeShellScriptAssignmentClient: %+v", err)
	}

	return &DeviceCustomAttributeShellScriptAssignmentClient{
		Client: client,
	}, nil
}
