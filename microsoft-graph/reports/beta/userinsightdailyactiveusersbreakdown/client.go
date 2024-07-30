package userinsightdailyactiveusersbreakdown

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightDailyActiveUsersBreakdownClient struct {
	Client *msgraph.Client
}

func NewUserInsightDailyActiveUsersBreakdownClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightDailyActiveUsersBreakdownClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightdailyactiveusersbreakdown", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightDailyActiveUsersBreakdownClient: %+v", err)
	}

	return &UserInsightDailyActiveUsersBreakdownClient{
		Client: client,
	}, nil
}
