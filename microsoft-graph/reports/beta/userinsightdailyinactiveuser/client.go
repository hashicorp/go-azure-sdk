package userinsightdailyinactiveuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightDailyInactiveUserClient struct {
	Client *msgraph.Client
}

func NewUserInsightDailyInactiveUserClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightDailyInactiveUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightdailyinactiveuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightDailyInactiveUserClient: %+v", err)
	}

	return &UserInsightDailyInactiveUserClient{
		Client: client,
	}, nil
}
