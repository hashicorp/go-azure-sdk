package userinsightmonthlyinactiveusersbyapplication

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightMonthlyInactiveUsersByApplicationClient struct {
	Client *msgraph.Client
}

func NewUserInsightMonthlyInactiveUsersByApplicationClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightMonthlyInactiveUsersByApplicationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightmonthlyinactiveusersbyapplication", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightMonthlyInactiveUsersByApplicationClient: %+v", err)
	}

	return &UserInsightMonthlyInactiveUsersByApplicationClient{
		Client: client,
	}, nil
}
