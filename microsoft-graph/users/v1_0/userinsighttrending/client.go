package userinsighttrending

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightTrendingClient struct {
	Client *msgraph.Client
}

func NewUserInsightTrendingClientWithBaseURI(api sdkEnv.Api) (*UserInsightTrendingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinsighttrending", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightTrendingClient: %+v", err)
	}

	return &UserInsightTrendingClient{
		Client: client,
	}, nil
}
