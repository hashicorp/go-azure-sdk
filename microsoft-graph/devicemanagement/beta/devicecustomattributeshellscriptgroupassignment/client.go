package devicecustomattributeshellscriptgroupassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeShellScriptGroupAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceCustomAttributeShellScriptGroupAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCustomAttributeShellScriptGroupAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecustomattributeshellscriptgroupassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCustomAttributeShellScriptGroupAssignmentClient: %+v", err)
	}

	return &DeviceCustomAttributeShellScriptGroupAssignmentClient{
		Client: client,
	}, nil
}
