package guestconfigurations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GuestconfigurationsClient struct {
	Client *resourcemanager.Client
}

func NewGuestconfigurationsClientWithBaseURI(sdkApi sdkEnv.Api) (*GuestconfigurationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "guestconfigurations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GuestconfigurationsClient: %+v", err)
	}

	return &GuestconfigurationsClient{
		Client: client,
	}, nil
}
