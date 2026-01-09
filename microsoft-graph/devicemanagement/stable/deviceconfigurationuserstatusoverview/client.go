package deviceconfigurationuserstatusoverview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationUserStatusOverviewClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationUserStatusOverviewClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationUserStatusOverviewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationuserstatusoverview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationUserStatusOverviewClient: %+v", err)
	}

	return &DeviceConfigurationUserStatusOverviewClient{
		Client: client,
	}, nil
}
