package v2023_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2023-01-01/ascusages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2023-01-01/caches"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2023-01-01/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2023-01-01/storagetargets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2023-01-01/usagemodels"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AscUsages      *ascusages.AscUsagesClient
	Caches         *caches.CachesClient
	SKUs           *skus.SKUsClient
	StorageTargets *storagetargets.StorageTargetsClient
	UsageModels    *usagemodels.UsageModelsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	ascUsagesClient, err := ascusages.NewAscUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AscUsages client: %+v", err)
	}
	configureFunc(ascUsagesClient.Client)

	cachesClient, err := caches.NewCachesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Caches client: %+v", err)
	}
	configureFunc(cachesClient.Client)

	sKUsClient, err := skus.NewSKUsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SKUs client: %+v", err)
	}
	configureFunc(sKUsClient.Client)

	storageTargetsClient, err := storagetargets.NewStorageTargetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StorageTargets client: %+v", err)
	}
	configureFunc(storageTargetsClient.Client)

	usageModelsClient, err := usagemodels.NewUsageModelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UsageModels client: %+v", err)
	}
	configureFunc(usageModelsClient.Client)

	return &Client{
		AscUsages:      ascUsagesClient,
		Caches:         cachesClient,
		SKUs:           sKUsClient,
		StorageTargets: storageTargetsClient,
		UsageModels:    usageModelsClient,
	}, nil
}
