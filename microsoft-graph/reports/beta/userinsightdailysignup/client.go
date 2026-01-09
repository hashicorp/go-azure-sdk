package userinsightdailysignup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightDailySignUpClient struct {
	Client *msgraph.Client
}

func NewUserInsightDailySignUpClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightDailySignUpClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightdailysignup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightDailySignUpClient: %+v", err)
	}

	return &UserInsightDailySignUpClient{
		Client: client,
	}, nil
}
