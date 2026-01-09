package userinsightmonthlyactiveuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightMonthlyActiveUserClient struct {
	Client *msgraph.Client
}

func NewUserInsightMonthlyActiveUserClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightMonthlyActiveUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightmonthlyactiveuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightMonthlyActiveUserClient: %+v", err)
	}

	return &UserInsightMonthlyActiveUserClient{
		Client: client,
	}, nil
}
