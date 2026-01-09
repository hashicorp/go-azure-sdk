package insightused

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightUsedClient struct {
	Client *msgraph.Client
}

func NewInsightUsedClientWithBaseURI(sdkApi sdkEnv.Api) (*InsightUsedClient, error) {
	client, err := msgraph.NewClient(sdkApi, "insightused", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InsightUsedClient: %+v", err)
	}

	return &InsightUsedClient{
		Client: client,
	}, nil
}
