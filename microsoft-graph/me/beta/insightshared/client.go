package insightshared

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightSharedClient struct {
	Client *msgraph.Client
}

func NewInsightSharedClientWithBaseURI(sdkApi sdkEnv.Api) (*InsightSharedClient, error) {
	client, err := msgraph.NewClient(sdkApi, "insightshared", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InsightSharedClient: %+v", err)
	}

	return &InsightSharedClient{
		Client: client,
	}, nil
}
