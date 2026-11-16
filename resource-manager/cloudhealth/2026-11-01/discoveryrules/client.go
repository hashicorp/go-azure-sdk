package discoveryrules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveryRulesClient struct {
	Client *resourcemanager.Client
}

func NewDiscoveryRulesClientWithBaseURI(sdkApi sdkEnv.Api) (*DiscoveryRulesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "discoveryrules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DiscoveryRulesClient: %+v", err)
	}

	return &DiscoveryRulesClient{
		Client: client,
	}, nil
}
