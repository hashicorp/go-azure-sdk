package deviceconfigurationdevicestatus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationDeviceStatusClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationDeviceStatusClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationDeviceStatusClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationdevicestatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationDeviceStatusClient: %+v", err)
	}

	return &DeviceConfigurationDeviceStatusClient{
		Client: client,
	}, nil
}
