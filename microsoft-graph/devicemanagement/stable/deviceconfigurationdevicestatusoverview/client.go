package deviceconfigurationdevicestatusoverview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationDeviceStatusOverviewClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationDeviceStatusOverviewClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationDeviceStatusOverviewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationdevicestatusoverview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationDeviceStatusOverviewClient: %+v", err)
	}

	return &DeviceConfigurationDeviceStatusOverviewClient{
		Client: client,
	}, nil
}
