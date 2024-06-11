package v2022_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resourcegraph/2022-10-01/graphqueries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resourcegraph/2022-10-01/graphquery"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resourcegraph/2022-10-01/resources"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	GraphQuery   *graphquery.GraphQueryClient
	Graphqueries *graphqueries.GraphqueriesClient
	Resources    *resources.ResourcesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	graphQueryClient, err := graphquery.NewGraphQueryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GraphQuery client: %+v", err)
	}
	configureFunc(graphQueryClient.Client)

	graphqueriesClient, err := graphqueries.NewGraphqueriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Graphqueries client: %+v", err)
	}
	configureFunc(graphqueriesClient.Client)

	resourcesClient, err := resources.NewResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Resources client: %+v", err)
	}
	configureFunc(resourcesClient.Client)

	return &Client{
		GraphQuery:   graphQueryClient,
		Graphqueries: graphqueriesClient,
		Resources:    resourcesClient,
	}, nil
}
