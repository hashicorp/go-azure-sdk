package v2022_04_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-04-01/openshiftclusters"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	OpenShiftClusters *openshiftclusters.OpenShiftClustersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	openShiftClustersClient := openshiftclusters.NewOpenShiftClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&openShiftClustersClient.Client)

	return Client{
		OpenShiftClusters: &openShiftClustersClient,
	}
}
