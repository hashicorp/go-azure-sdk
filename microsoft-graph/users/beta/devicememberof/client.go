package devicememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceMemberOfClient struct {
	Client *msgraph.Client
}

func NewDeviceMemberOfClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceMemberOfClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceMemberOfClient: %+v", err)
	}

	return &DeviceMemberOfClient{
		Client: client,
	}, nil
}
