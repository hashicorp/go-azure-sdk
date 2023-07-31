package generaterecommendations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenerateRecommendationsClient struct {
	Client *resourcemanager.Client
}

func NewGenerateRecommendationsClientWithBaseURI(api environments.Api) (*GenerateRecommendationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "generaterecommendations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GenerateRecommendationsClient: %+v", err)
	}

	return &GenerateRecommendationsClient{
		Client: client,
	}, nil
}
