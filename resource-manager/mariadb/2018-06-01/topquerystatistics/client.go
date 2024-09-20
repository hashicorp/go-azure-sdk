package topquerystatistics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TopQueryStatisticsClient struct {
	Client *resourcemanager.Client
}

func NewTopQueryStatisticsClientWithBaseURI(sdkApi sdkEnv.Api) (*TopQueryStatisticsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "topquerystatistics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TopQueryStatisticsClient: %+v", err)
	}

	return &TopQueryStatisticsClient{
		Client: client,
	}, nil
}
