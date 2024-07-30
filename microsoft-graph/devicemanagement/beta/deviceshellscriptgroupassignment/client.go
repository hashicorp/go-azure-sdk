package deviceshellscriptgroupassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceShellScriptGroupAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceShellScriptGroupAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceShellScriptGroupAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceshellscriptgroupassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceShellScriptGroupAssignmentClient: %+v", err)
	}

	return &DeviceShellScriptGroupAssignmentClient{
		Client: client,
	}, nil
}
