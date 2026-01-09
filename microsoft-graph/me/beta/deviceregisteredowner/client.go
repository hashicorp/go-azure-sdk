package deviceregisteredowner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceRegisteredOwnerClient struct {
	Client *msgraph.Client
}

func NewDeviceRegisteredOwnerClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceRegisteredOwnerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceregisteredowner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceRegisteredOwnerClient: %+v", err)
	}

	return &DeviceRegisteredOwnerClient{
		Client: client,
	}, nil
}
