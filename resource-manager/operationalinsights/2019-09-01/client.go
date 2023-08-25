package v2019_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2019-09-01/querypackqueries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2019-09-01/querypacks"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	QueryPackQueries *querypackqueries.QueryPackQueriesClient
	QueryPacks       *querypacks.QueryPacksClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	queryPackQueriesClient, err := querypackqueries.NewQueryPackQueriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QueryPackQueries client: %+v", err)
	}
	configureFunc(queryPackQueriesClient.Client)

	queryPacksClient, err := querypacks.NewQueryPacksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QueryPacks client: %+v", err)
	}
	configureFunc(queryPacksClient.Client)

	return &Client{
		QueryPackQueries: queryPackQueriesClient,
		QueryPacks:       queryPacksClient,
	}, nil
}
