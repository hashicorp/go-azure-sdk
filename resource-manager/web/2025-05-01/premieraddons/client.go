package premieraddons

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PremierAddonsClient struct {
	Client *resourcemanager.Client
}

func NewPremierAddonsClientWithBaseURI(sdkApi sdkEnv.Api) (*PremierAddonsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "premieraddons", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PremierAddonsClient: %+v", err)
	}

	return &PremierAddonsClient{
		Client: client,
	}, nil
}
