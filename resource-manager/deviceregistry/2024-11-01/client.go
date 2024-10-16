package v2024_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-11-01/assetendpointprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-11-01/assets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-11-01/billingcontainers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AssetEndpointProfiles *assetendpointprofiles.AssetEndpointProfilesClient
	Assets                *assets.AssetsClient
	BillingContainers     *billingcontainers.BillingContainersClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	assetEndpointProfilesClient, err := assetendpointprofiles.NewAssetEndpointProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AssetEndpointProfiles client: %+v", err)
	}
	configureFunc(assetEndpointProfilesClient.Client)

	assetsClient, err := assets.NewAssetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Assets client: %+v", err)
	}
	configureFunc(assetsClient.Client)

	billingContainersClient, err := billingcontainers.NewBillingContainersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingContainers client: %+v", err)
	}
	configureFunc(billingContainersClient.Client)

	return &Client{
		AssetEndpointProfiles: assetEndpointProfilesClient,
		Assets:                assetsClient,
		BillingContainers:     billingContainersClient,
	}, nil
}
