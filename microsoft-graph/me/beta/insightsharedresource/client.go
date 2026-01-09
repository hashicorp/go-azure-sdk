package insightsharedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightSharedResourceClient struct {
	Client *msgraph.Client
}

func NewInsightSharedResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*InsightSharedResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "insightsharedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InsightSharedResourceClient: %+v", err)
	}

	return &InsightSharedResourceClient{
		Client: client,
	}, nil
}
