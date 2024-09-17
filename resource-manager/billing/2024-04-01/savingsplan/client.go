package savingsplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SavingsPlanClient struct {
	Client *resourcemanager.Client
}

func NewSavingsPlanClientWithBaseURI(sdkApi sdkEnv.Api) (*SavingsPlanClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "savingsplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SavingsPlanClient: %+v", err)
	}

	return &SavingsPlanClient{
		Client: client,
	}, nil
}
