package recommendations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationsClient struct {
	Client *resourcemanager.Client
}

func NewRecommendationsClientWithBaseURI(api environments.Api) (*RecommendationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "recommendations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RecommendationsClient: %+v", err)
	}

	return &RecommendationsClient{
		Client: client,
	}, nil
}
