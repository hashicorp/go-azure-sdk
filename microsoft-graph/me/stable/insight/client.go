package insight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightClient struct {
	Client *msgraph.Client
}

func NewInsightClientWithBaseURI(sdkApi sdkEnv.Api) (*InsightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "insight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InsightClient: %+v", err)
	}

	return &InsightClient{
		Client: client,
	}, nil
}
