package v2018_09_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/customproviders/2018-09-01-preview/associations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/customproviders/2018-09-01-preview/customresourceprovider"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Associations           *associations.AssociationsClient
	CustomResourceProvider *customresourceprovider.CustomResourceProviderClient
}

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	associationsClient, err := associations.NewAssociationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Associations client: %+v", err)
	}
	configureFunc(associationsClient.Client)

	customResourceProviderClient, err := customresourceprovider.NewCustomResourceProviderClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CustomResourceProvider client: %+v", err)
	}
	configureFunc(customResourceProviderClient.Client)

	return &Client{
		Associations:           associationsClient,
		CustomResourceProvider: customResourceProviderClient,
	}, nil
}
