package deviceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceExtensionClient struct {
	Client *msgraph.Client
}

func NewDeviceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceExtensionClient: %+v", err)
	}

	return &DeviceExtensionClient{
		Client: client,
	}, nil
}
