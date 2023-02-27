// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2019_09_01

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
