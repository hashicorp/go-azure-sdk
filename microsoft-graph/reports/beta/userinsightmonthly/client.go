package userinsightmonthly

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightMonthlyClient struct {
	Client *msgraph.Client
}

func NewUserInsightMonthlyClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightMonthlyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightmonthly", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightMonthlyClient: %+v", err)
	}

	return &UserInsightMonthlyClient{
		Client: client,
	}, nil
}
