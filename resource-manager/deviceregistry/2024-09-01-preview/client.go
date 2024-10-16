package v2024_09_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/assetendpointprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/assets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/billingcontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/discoveredassetendpointprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/discoveredassets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/schemaregistries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/schemas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/schemaversions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AssetEndpointProfiles           *assetendpointprofiles.AssetEndpointProfilesClient
	Assets                          *assets.AssetsClient
	BillingContainers               *billingcontainers.BillingContainersClient
	DiscoveredAssetEndpointProfiles *discoveredassetendpointprofiles.DiscoveredAssetEndpointProfilesClient
	DiscoveredAssets                *discoveredassets.DiscoveredAssetsClient
	SchemaRegistries                *schemaregistries.SchemaRegistriesClient
	SchemaVersions                  *schemaversions.SchemaVersionsClient
	Schemas                         *schemas.SchemasClient
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

	discoveredAssetEndpointProfilesClient, err := discoveredassetendpointprofiles.NewDiscoveredAssetEndpointProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiscoveredAssetEndpointProfiles client: %+v", err)
	}
	configureFunc(discoveredAssetEndpointProfilesClient.Client)

	discoveredAssetsClient, err := discoveredassets.NewDiscoveredAssetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiscoveredAssets client: %+v", err)
	}
	configureFunc(discoveredAssetsClient.Client)

	schemaRegistriesClient, err := schemaregistries.NewSchemaRegistriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SchemaRegistries client: %+v", err)
	}
	configureFunc(schemaRegistriesClient.Client)

	schemaVersionsClient, err := schemaversions.NewSchemaVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SchemaVersions client: %+v", err)
	}
	configureFunc(schemaVersionsClient.Client)

	schemasClient, err := schemas.NewSchemasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Schemas client: %+v", err)
	}
	configureFunc(schemasClient.Client)

	return &Client{
		AssetEndpointProfiles:           assetEndpointProfilesClient,
		Assets:                          assetsClient,
		BillingContainers:               billingContainersClient,
		DiscoveredAssetEndpointProfiles: discoveredAssetEndpointProfilesClient,
		DiscoveredAssets:                discoveredAssetsClient,
		SchemaRegistries:                schemaRegistriesClient,
		SchemaVersions:                  schemaVersionsClient,
		Schemas:                         schemasClient,
	}, nil
}
