package recommendationconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationConfigurationClient struct {
	Client *msgraph.Client
}

func NewRecommendationConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*RecommendationConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "recommendationconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RecommendationConfigurationClient: %+v", err)
	}

	return &RecommendationConfigurationClient{
		Client: client,
	}, nil
}
