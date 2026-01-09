package deviceconfigurationuserstatus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationUserStatusClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationUserStatusClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationUserStatusClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationuserstatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationUserStatusClient: %+v", err)
	}

	return &DeviceConfigurationUserStatusClient{
		Client: client,
	}, nil
}
