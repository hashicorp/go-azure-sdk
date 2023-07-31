package networkrulesets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkRuleSetsClient struct {
	Client *resourcemanager.Client
}

func NewNetworkRuleSetsClientWithBaseURI(api sdkEnv.Api) (*NetworkRuleSetsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "networkrulesets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetworkRuleSetsClient: %+v", err)
	}

	return &NetworkRuleSetsClient{
		Client: client,
	}, nil
}
