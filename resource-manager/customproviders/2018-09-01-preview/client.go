package v2018_09_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/customproviders/2018-09-01-preview/associations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/customproviders/2018-09-01-preview/customresourceprovider"
)

type Client struct {
	Associations           *associations.AssociationsClient
	CustomResourceProvider *customresourceprovider.CustomResourceProviderClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	associationsClient := associations.NewAssociationsClientWithBaseURI(endpoint)
	configureAuthFunc(&associationsClient.Client)

	customResourceProviderClient := customresourceprovider.NewCustomResourceProviderClientWithBaseURI(endpoint)
	configureAuthFunc(&customResourceProviderClient.Client)

	return Client{
		Associations:           &associationsClient,
		CustomResourceProvider: &customResourceProviderClient,
	}
}
