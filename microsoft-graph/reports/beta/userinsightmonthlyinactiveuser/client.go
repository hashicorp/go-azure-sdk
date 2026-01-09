package userinsightmonthlyinactiveuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightMonthlyInactiveUserClient struct {
	Client *msgraph.Client
}

func NewUserInsightMonthlyInactiveUserClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightMonthlyInactiveUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightmonthlyinactiveuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightMonthlyInactiveUserClient: %+v", err)
	}

	return &UserInsightMonthlyInactiveUserClient{
		Client: client,
	}, nil
}
