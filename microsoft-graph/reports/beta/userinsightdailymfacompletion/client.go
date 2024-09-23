package userinsightdailymfacompletion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightDailyMfaCompletionClient struct {
	Client *msgraph.Client
}

func NewUserInsightDailyMfaCompletionClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightDailyMfaCompletionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightdailymfacompletion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightDailyMfaCompletionClient: %+v", err)
	}

	return &UserInsightDailyMfaCompletionClient{
		Client: client,
	}, nil
}
