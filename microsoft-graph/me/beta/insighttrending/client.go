package insighttrending

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightTrendingClient struct {
	Client *msgraph.Client
}

func NewInsightTrendingClientWithBaseURI(sdkApi sdkEnv.Api) (*InsightTrendingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "insighttrending", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InsightTrendingClient: %+v", err)
	}

	return &InsightTrendingClient{
		Client: client,
	}, nil
}
