package v2025_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/assetendpointprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/assets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/billingcontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/namespaceassets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/namespacedevices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/namespacediscoveredassets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/namespacediscovereddevices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/schemaregistries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/schemas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/schemaversions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AssetEndpointProfiles      *assetendpointprofiles.AssetEndpointProfilesClient
	Assets                     *assets.AssetsClient
	BillingContainers          *billingcontainers.BillingContainersClient
	NamespaceAssets            *namespaceassets.NamespaceAssetsClient
	NamespaceDevices           *namespacedevices.NamespaceDevicesClient
	NamespaceDiscoveredAssets  *namespacediscoveredassets.NamespaceDiscoveredAssetsClient
	NamespaceDiscoveredDevices *namespacediscovereddevices.NamespaceDiscoveredDevicesClient
	Namespaces                 *namespaces.NamespacesClient
	SchemaRegistries           *schemaregistries.SchemaRegistriesClient
	SchemaVersions             *schemaversions.SchemaVersionsClient
	Schemas                    *schemas.SchemasClient
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

	namespaceAssetsClient, err := namespaceassets.NewNamespaceAssetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NamespaceAssets client: %+v", err)
	}
	configureFunc(namespaceAssetsClient.Client)

	namespaceDevicesClient, err := namespacedevices.NewNamespaceDevicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NamespaceDevices client: %+v", err)
	}
	configureFunc(namespaceDevicesClient.Client)

	namespaceDiscoveredAssetsClient, err := namespacediscoveredassets.NewNamespaceDiscoveredAssetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NamespaceDiscoveredAssets client: %+v", err)
	}
	configureFunc(namespaceDiscoveredAssetsClient.Client)

	namespaceDiscoveredDevicesClient, err := namespacediscovereddevices.NewNamespaceDiscoveredDevicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NamespaceDiscoveredDevices client: %+v", err)
	}
	configureFunc(namespaceDiscoveredDevicesClient.Client)

	namespacesClient, err := namespaces.NewNamespacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Namespaces client: %+v", err)
	}
	configureFunc(namespacesClient.Client)

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
		AssetEndpointProfiles:      assetEndpointProfilesClient,
		Assets:                     assetsClient,
		BillingContainers:          billingContainersClient,
		NamespaceAssets:            namespaceAssetsClient,
		NamespaceDevices:           namespaceDevicesClient,
		NamespaceDiscoveredAssets:  namespaceDiscoveredAssetsClient,
		NamespaceDiscoveredDevices: namespaceDiscoveredDevicesClient,
		Namespaces:                 namespacesClient,
		SchemaRegistries:           schemaRegistriesClient,
		SchemaVersions:             schemaVersionsClient,
		Schemas:                    schemasClient,
	}, nil
}
