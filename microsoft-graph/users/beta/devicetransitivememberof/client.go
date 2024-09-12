package devicetransitivememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceTransitiveMemberOfClient struct {
	Client *msgraph.Client
}

func NewDeviceTransitiveMemberOfClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceTransitiveMemberOfClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicetransitivememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceTransitiveMemberOfClient: %+v", err)
	}

	return &DeviceTransitiveMemberOfClient{
		Client: client,
	}, nil
}
