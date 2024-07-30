package userinsightdailyinactiveusersbyapplication

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightDailyInactiveUsersByApplicationClient struct {
	Client *msgraph.Client
}

func NewUserInsightDailyInactiveUsersByApplicationClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightDailyInactiveUsersByApplicationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightdailyinactiveusersbyapplication", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightDailyInactiveUsersByApplicationClient: %+v", err)
	}

	return &UserInsightDailyInactiveUsersByApplicationClient{
		Client: client,
	}, nil
}
