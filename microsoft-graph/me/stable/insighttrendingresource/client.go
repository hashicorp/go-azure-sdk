package insighttrendingresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightTrendingResourceClient struct {
	Client *msgraph.Client
}

func NewInsightTrendingResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*InsightTrendingResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "insighttrendingresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InsightTrendingResourceClient: %+v", err)
	}

	return &InsightTrendingResourceClient{
		Client: client,
	}, nil
}
