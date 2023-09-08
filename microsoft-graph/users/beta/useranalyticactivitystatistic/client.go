package useranalyticactivitystatistic

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAnalyticActivityStatisticClient struct {
	Client *msgraph.Client
}

func NewUserAnalyticActivityStatisticClientWithBaseURI(api sdkEnv.Api) (*UserAnalyticActivityStatisticClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useranalyticactivitystatistic", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAnalyticActivityStatisticClient: %+v", err)
	}

	return &UserAnalyticActivityStatisticClient{
		Client: client,
	}, nil
}
