package addons

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddonsClient struct {
	Client *resourcemanager.Client
}

func NewAddonsClientWithBaseURI(sdkApi sdkEnv.Api) (*AddonsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "addons", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AddonsClient: %+v", err)
	}

	return &AddonsClient{
		Client: client,
	}, nil
}
