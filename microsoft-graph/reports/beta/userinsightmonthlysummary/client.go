package userinsightmonthlysummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightMonthlySummaryClient struct {
	Client *msgraph.Client
}

func NewUserInsightMonthlySummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightMonthlySummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightmonthlysummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightMonthlySummaryClient: %+v", err)
	}

	return &UserInsightMonthlySummaryClient{
		Client: client,
	}, nil
}
