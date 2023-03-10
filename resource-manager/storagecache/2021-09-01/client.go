package v2021_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2021-09-01/caches"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2021-09-01/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2021-09-01/storagetargets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2021-09-01/usagemodels"
)

type Client struct {
	Caches         *caches.CachesClient
	SKUs           *skus.SKUsClient
	StorageTargets *storagetargets.StorageTargetsClient
	UsageModels    *usagemodels.UsageModelsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	cachesClient := caches.NewCachesClientWithBaseURI(endpoint)
	configureAuthFunc(&cachesClient.Client)

	sKUsClient := skus.NewSKUsClientWithBaseURI(endpoint)
	configureAuthFunc(&sKUsClient.Client)

	storageTargetsClient := storagetargets.NewStorageTargetsClientWithBaseURI(endpoint)
	configureAuthFunc(&storageTargetsClient.Client)

	usageModelsClient := usagemodels.NewUsageModelsClientWithBaseURI(endpoint)
	configureAuthFunc(&usageModelsClient.Client)

	return Client{
		Caches:         &cachesClient,
		SKUs:           &sKUsClient,
		StorageTargets: &storageTargetsClient,
		UsageModels:    &usageModelsClient,
	}
}
