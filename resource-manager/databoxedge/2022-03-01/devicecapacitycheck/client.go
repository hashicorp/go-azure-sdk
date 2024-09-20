package devicecapacitycheck

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCapacityCheckClient struct {
	Client *resourcemanager.Client
}

func NewDeviceCapacityCheckClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCapacityCheckClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "devicecapacitycheck", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCapacityCheckClient: %+v", err)
	}

	return &DeviceCapacityCheckClient{
		Client: client,
	}, nil
}
