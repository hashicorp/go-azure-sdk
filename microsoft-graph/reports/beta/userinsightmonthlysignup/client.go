package userinsightmonthlysignup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightMonthlySignUpClient struct {
	Client *msgraph.Client
}

func NewUserInsightMonthlySignUpClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightMonthlySignUpClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightmonthlysignup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightMonthlySignUpClient: %+v", err)
	}

	return &UserInsightMonthlySignUpClient{
		Client: client,
	}, nil
}
