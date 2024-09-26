package customrecommendations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomRecommendationsClient struct {
	Client *resourcemanager.Client
}

func NewCustomRecommendationsClientWithBaseURI(sdkApi sdkEnv.Api) (*CustomRecommendationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "customrecommendations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomRecommendationsClient: %+v", err)
	}

	return &CustomRecommendationsClient{
		Client: client,
	}, nil
}
