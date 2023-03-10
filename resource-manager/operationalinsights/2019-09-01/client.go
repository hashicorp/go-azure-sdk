package v2019_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2019-09-01/querypackqueries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2019-09-01/querypacks"
)

type Client struct {
	QueryPackQueries *querypackqueries.QueryPackQueriesClient
	QueryPacks       *querypacks.QueryPacksClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	queryPackQueriesClient := querypackqueries.NewQueryPackQueriesClientWithBaseURI(endpoint)
	configureAuthFunc(&queryPackQueriesClient.Client)

	queryPacksClient := querypacks.NewQueryPacksClientWithBaseURI(endpoint)
	configureAuthFunc(&queryPacksClient.Client)

	return Client{
		QueryPackQueries: &queryPackQueriesClient,
		QueryPacks:       &queryPacksClient,
	}
}
