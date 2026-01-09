package insightusedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightUsedResourceClient struct {
	Client *msgraph.Client
}

func NewInsightUsedResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*InsightUsedResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "insightusedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InsightUsedResourceClient: %+v", err)
	}

	return &InsightUsedResourceClient{
		Client: client,
	}, nil
}
