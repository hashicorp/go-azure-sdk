package recommendedactions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionsClient struct {
	Client *resourcemanager.Client
}

func NewRecommendedActionsClientWithBaseURI(sdkApi sdkEnv.Api) (*RecommendedActionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "recommendedactions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RecommendedActionsClient: %+v", err)
	}

	return &RecommendedActionsClient{
		Client: client,
	}, nil
}
