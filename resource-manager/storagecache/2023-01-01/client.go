// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2023_01_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2023-01-01/ascusages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2023-01-01/caches"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2023-01-01/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2023-01-01/storagetargets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2023-01-01/usagemodels"
)

type Client struct {
	AscUsages      *ascusages.AscUsagesClient
	Caches         *caches.CachesClient
	SKUs           *skus.SKUsClient
	StorageTargets *storagetargets.StorageTargetsClient
	UsageModels    *usagemodels.UsageModelsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	ascUsagesClient := ascusages.NewAscUsagesClientWithBaseURI(endpoint)
	configureAuthFunc(&ascUsagesClient.Client)

	cachesClient := caches.NewCachesClientWithBaseURI(endpoint)
	configureAuthFunc(&cachesClient.Client)

	sKUsClient := skus.NewSKUsClientWithBaseURI(endpoint)
	configureAuthFunc(&sKUsClient.Client)

	storageTargetsClient := storagetargets.NewStorageTargetsClientWithBaseURI(endpoint)
	configureAuthFunc(&storageTargetsClient.Client)

	usageModelsClient := usagemodels.NewUsageModelsClientWithBaseURI(endpoint)
	configureAuthFunc(&usageModelsClient.Client)

	return Client{
		AscUsages:      &ascUsagesClient,
		Caches:         &cachesClient,
		SKUs:           &sKUsClient,
		StorageTargets: &storageTargetsClient,
		UsageModels:    &usageModelsClient,
	}
}
