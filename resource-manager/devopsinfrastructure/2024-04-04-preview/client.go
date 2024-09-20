package v2024_04_04_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/devopsinfrastructure/2024-04-04-preview/imageversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devopsinfrastructure/2024-04-04-preview/pools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devopsinfrastructure/2024-04-04-preview/resourcedetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devopsinfrastructure/2024-04-04-preview/sku"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devopsinfrastructure/2024-04-04-preview/subscriptionusages"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ImageVersions      *imageversions.ImageVersionsClient
	Pools              *pools.PoolsClient
	ResourceDetails    *resourcedetails.ResourceDetailsClient
	Sku                *sku.SkuClient
	SubscriptionUsages *subscriptionusages.SubscriptionUsagesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	imageVersionsClient, err := imageversions.NewImageVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ImageVersions client: %+v", err)
	}
	configureFunc(imageVersionsClient.Client)

	poolsClient, err := pools.NewPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Pools client: %+v", err)
	}
	configureFunc(poolsClient.Client)

	resourceDetailsClient, err := resourcedetails.NewResourceDetailsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceDetails client: %+v", err)
	}
	configureFunc(resourceDetailsClient.Client)

	skuClient, err := sku.NewSkuClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Sku client: %+v", err)
	}
	configureFunc(skuClient.Client)

	subscriptionUsagesClient, err := subscriptionusages.NewSubscriptionUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubscriptionUsages client: %+v", err)
	}
	configureFunc(subscriptionUsagesClient.Client)

	return &Client{
		ImageVersions:      imageVersionsClient,
		Pools:              poolsClient,
		ResourceDetails:    resourceDetailsClient,
		Sku:                skuClient,
		SubscriptionUsages: subscriptionUsagesClient,
	}, nil
}
