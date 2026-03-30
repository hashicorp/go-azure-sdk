package l2networks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type L2NetworksClient struct {
	Client *resourcemanager.Client
}

func NewL2NetworksClientWithBaseURI(sdkApi sdkEnv.Api) (*L2NetworksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "l2networks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating L2NetworksClient: %+v", err)
	}

	return &L2NetworksClient{
		Client: client,
	}, nil
}
