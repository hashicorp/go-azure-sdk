package racks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RacksClient struct {
	Client *resourcemanager.Client
}

func NewRacksClientWithBaseURI(sdkApi sdkEnv.Api) (*RacksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "racks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RacksClient: %+v", err)
	}

	return &RacksClient{
		Client: client,
	}, nil
}
