package userinsightmonthlymfacompletion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightMonthlyMfaCompletionClient struct {
	Client *msgraph.Client
}

func NewUserInsightMonthlyMfaCompletionClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightMonthlyMfaCompletionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightmonthlymfacompletion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightMonthlyMfaCompletionClient: %+v", err)
	}

	return &UserInsightMonthlyMfaCompletionClient{
		Client: client,
	}, nil
}
