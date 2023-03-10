package v2022_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2022-05-01/links"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2022-05-01/servicelinker"
)

type Client struct {
	Links         *links.LinksClient
	ServiceLinker *servicelinker.ServiceLinkerClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	linksClient := links.NewLinksClientWithBaseURI(endpoint)
	configureAuthFunc(&linksClient.Client)

	serviceLinkerClient := servicelinker.NewServiceLinkerClientWithBaseURI(endpoint)
	configureAuthFunc(&serviceLinkerClient.Client)

	return Client{
		Links:         &linksClient,
		ServiceLinker: &serviceLinkerClient,
	}
}
