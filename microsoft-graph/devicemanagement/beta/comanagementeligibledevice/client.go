package comanagementeligibledevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceClient struct {
	Client *msgraph.Client
}

func NewComanagementEligibleDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagementEligibleDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanagementeligibledevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagementEligibleDeviceClient: %+v", err)
	}

	return &ComanagementEligibleDeviceClient{
		Client: client,
	}, nil
}
