package devicemanagementpartner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPartnerClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementPartnerClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementPartnerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementpartner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementPartnerClient: %+v", err)
	}

	return &DeviceManagementPartnerClient{
		Client: client,
	}, nil
}
