package devicemanagement

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagement", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementClient: %+v", err)
	}

	return &DeviceManagementClient{
		Client: client,
	}, nil
}
