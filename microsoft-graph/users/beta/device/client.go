package device

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceClient struct {
	Client *msgraph.Client
}

func NewDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "device", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceClient: %+v", err)
	}

	return &DeviceClient{
		Client: client,
	}, nil
}
