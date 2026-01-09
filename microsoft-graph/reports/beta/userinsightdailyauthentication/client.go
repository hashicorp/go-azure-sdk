package userinsightdailyauthentication

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightDailyAuthenticationClient struct {
	Client *msgraph.Client
}

func NewUserInsightDailyAuthenticationClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightDailyAuthenticationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightdailyauthentication", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightDailyAuthenticationClient: %+v", err)
	}

	return &UserInsightDailyAuthenticationClient{
		Client: client,
	}, nil
}
