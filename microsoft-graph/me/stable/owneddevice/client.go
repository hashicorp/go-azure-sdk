package owneddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OwnedDeviceClient struct {
	Client *msgraph.Client
}

func NewOwnedDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*OwnedDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "owneddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OwnedDeviceClient: %+v", err)
	}

	return &OwnedDeviceClient{
		Client: client,
	}, nil
}
