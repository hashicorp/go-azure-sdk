package costdetails

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostDetailsClient struct {
	Client *resourcemanager.Client
}

func NewCostDetailsClientWithBaseURI(sdkApi sdkEnv.Api) (*CostDetailsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "costdetails", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CostDetailsClient: %+v", err)
	}

	return &CostDetailsClient{
		Client: client,
	}, nil
}
