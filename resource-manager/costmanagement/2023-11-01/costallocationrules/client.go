package costallocationrules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostAllocationRulesClient struct {
	Client *resourcemanager.Client
}

func NewCostAllocationRulesClientWithBaseURI(sdkApi sdkEnv.Api) (*CostAllocationRulesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "costallocationrules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CostAllocationRulesClient: %+v", err)
	}

	return &CostAllocationRulesClient{
		Client: client,
	}, nil
}
