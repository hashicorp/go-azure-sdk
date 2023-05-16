package v2022_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-04-01/openshiftclusters"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	OpenShiftClusters *openshiftclusters.OpenShiftClustersClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	openShiftClustersClient, err := openshiftclusters.NewOpenShiftClustersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building OpenShiftClusters client: %+v", err)
	}
	configureFunc(openShiftClustersClient.Client)

	return &Client{
		OpenShiftClusters: openShiftClustersClient,
	}, nil
}
