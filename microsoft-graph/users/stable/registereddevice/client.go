package registereddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegisteredDeviceClient struct {
	Client *msgraph.Client
}

func NewRegisteredDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*RegisteredDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "registereddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RegisteredDeviceClient: %+v", err)
	}

	return &RegisteredDeviceClient{
		Client: client,
	}, nil
}
