package userinsightmonthlyauthentication

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightMonthlyAuthenticationClient struct {
	Client *msgraph.Client
}

func NewUserInsightMonthlyAuthenticationClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightMonthlyAuthenticationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightmonthlyauthentication", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightMonthlyAuthenticationClient: %+v", err)
	}

	return &UserInsightMonthlyAuthenticationClient{
		Client: client,
	}, nil
}
