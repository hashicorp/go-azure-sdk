package policysets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySetsClient struct {
	Client *resourcemanager.Client
}

func NewPolicySetsClientWithBaseURI(sdkApi sdkEnv.Api) (*PolicySetsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "policysets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicySetsClient: %+v", err)
	}

	return &PolicySetsClient{
		Client: client,
	}, nil
}
