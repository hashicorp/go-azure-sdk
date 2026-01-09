package userinsightdailysummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightDailySummaryClient struct {
	Client *msgraph.Client
}

func NewUserInsightDailySummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightDailySummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightdailysummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightDailySummaryClient: %+v", err)
	}

	return &UserInsightDailySummaryClient{
		Client: client,
	}, nil
}
