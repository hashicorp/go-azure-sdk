package deviceregistereduser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceRegisteredUserClient struct {
	Client *msgraph.Client
}

func NewDeviceRegisteredUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceRegisteredUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceregistereduser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceRegisteredUserClient: %+v", err)
	}

	return &DeviceRegisteredUserClient{
		Client: client,
	}, nil
}
