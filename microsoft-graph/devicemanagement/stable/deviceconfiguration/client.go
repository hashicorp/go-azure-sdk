package deviceconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationClient: %+v", err)
	}

	return &DeviceConfigurationClient{
		Client: client,
	}, nil
}
