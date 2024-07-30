package deviceusageright

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceUsageRightClient struct {
	Client *msgraph.Client
}

func NewDeviceUsageRightClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceUsageRightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceusageright", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceUsageRightClient: %+v", err)
	}

	return &DeviceUsageRightClient{
		Client: client,
	}, nil
}
