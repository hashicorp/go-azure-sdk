package recommendationrules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationRulesClient struct {
	Client *resourcemanager.Client
}

func NewRecommendationRulesClientWithBaseURI(sdkApi sdkEnv.Api) (*RecommendationRulesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "recommendationrules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RecommendationRulesClient: %+v", err)
	}

	return &RecommendationRulesClient{
		Client: client,
	}, nil
}
