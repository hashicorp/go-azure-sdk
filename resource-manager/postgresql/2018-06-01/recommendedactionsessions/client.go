package recommendedactionsessions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionSessionsClient struct {
	Client *resourcemanager.Client
}

func NewRecommendedActionSessionsClientWithBaseURI(sdkApi sdkEnv.Api) (*RecommendedActionSessionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "recommendedactionsessions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RecommendedActionSessionsClient: %+v", err)
	}

	return &RecommendedActionSessionsClient{
		Client: client,
	}, nil
}
