package userinsightdailyactiveuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightDailyActiveUserClient struct {
	Client *msgraph.Client
}

func NewUserInsightDailyActiveUserClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightDailyActiveUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightdailyactiveuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightDailyActiveUserClient: %+v", err)
	}

	return &UserInsightDailyActiveUserClient{
		Client: client,
	}, nil
}
