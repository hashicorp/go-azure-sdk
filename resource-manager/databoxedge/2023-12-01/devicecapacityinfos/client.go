package devicecapacityinfos

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCapacityInfosClient struct {
	Client *resourcemanager.Client
}

func NewDeviceCapacityInfosClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCapacityInfosClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "devicecapacityinfos", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCapacityInfosClient: %+v", err)
	}

	return &DeviceCapacityInfosClient{
		Client: client,
	}, nil
}
