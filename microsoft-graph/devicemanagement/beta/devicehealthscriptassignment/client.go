package devicehealthscriptassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceHealthScriptAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceHealthScriptAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicehealthscriptassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceHealthScriptAssignmentClient: %+v", err)
	}

	return &DeviceHealthScriptAssignmentClient{
		Client: client,
	}, nil
}
