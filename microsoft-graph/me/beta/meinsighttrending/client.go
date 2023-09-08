package meinsighttrending

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInsightTrendingClient struct {
	Client *msgraph.Client
}

func NewMeInsightTrendingClientWithBaseURI(api sdkEnv.Api) (*MeInsightTrendingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinsighttrending", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInsightTrendingClient: %+v", err)
	}

	return &MeInsightTrendingClient{
		Client: client,
	}, nil
}
