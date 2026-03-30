package l3networks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type L3NetworksClient struct {
	Client *resourcemanager.Client
}

func NewL3NetworksClientWithBaseURI(sdkApi sdkEnv.Api) (*L3NetworksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "l3networks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating L3NetworksClient: %+v", err)
	}

	return &L3NetworksClient{
		Client: client,
	}, nil
}
