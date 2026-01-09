package insightsharedlastsharedmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightSharedLastSharedMethodClient struct {
	Client *msgraph.Client
}

func NewInsightSharedLastSharedMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*InsightSharedLastSharedMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "insightsharedlastsharedmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InsightSharedLastSharedMethodClient: %+v", err)
	}

	return &InsightSharedLastSharedMethodClient{
		Client: client,
	}, nil
}
