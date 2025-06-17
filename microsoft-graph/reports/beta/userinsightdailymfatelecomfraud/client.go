package userinsightdailymfatelecomfraud

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightDailyMfaTelecomFraudClient struct {
	Client *msgraph.Client
}

func NewUserInsightDailyMfaTelecomFraudClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightDailyMfaTelecomFraudClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightdailymfatelecomfraud", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightDailyMfaTelecomFraudClient: %+v", err)
	}

	return &UserInsightDailyMfaTelecomFraudClient{
		Client: client,
	}, nil
}
