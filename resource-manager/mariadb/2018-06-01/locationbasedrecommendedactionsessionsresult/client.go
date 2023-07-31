package locationbasedrecommendedactionsessionsresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationBasedRecommendedActionSessionsResultClient struct {
	Client *resourcemanager.Client
}

func NewLocationBasedRecommendedActionSessionsResultClientWithBaseURI(api sdkEnv.Api) (*LocationBasedRecommendedActionSessionsResultClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "locationbasedrecommendedactionsessionsresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LocationBasedRecommendedActionSessionsResultClient: %+v", err)
	}

	return &LocationBasedRecommendedActionSessionsResultClient{
		Client: client,
	}, nil
}
