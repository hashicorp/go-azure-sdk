package v2021_11_20_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/elasticsan/2021-11-20-preview/elasticsan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elasticsan/2021-11-20-preview/elasticsans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elasticsan/2021-11-20-preview/elasticsanskus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elasticsan/2021-11-20-preview/volumegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elasticsan/2021-11-20-preview/volumes"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ElasticSan     *elasticsan.ElasticSanClient
	ElasticSanSkus *elasticsanskus.ElasticSanSkusClient
	ElasticSans    *elasticsans.ElasticSansClient
	VolumeGroups   *volumegroups.VolumeGroupsClient
	Volumes        *volumes.VolumesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	elasticSanClient, err := elasticsan.NewElasticSanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ElasticSan client: %+v", err)
	}
	configureFunc(elasticSanClient.Client)

	elasticSanSkusClient, err := elasticsanskus.NewElasticSanSkusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ElasticSanSkus client: %+v", err)
	}
	configureFunc(elasticSanSkusClient.Client)

	elasticSansClient, err := elasticsans.NewElasticSansClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ElasticSans client: %+v", err)
	}
	configureFunc(elasticSansClient.Client)

	volumeGroupsClient, err := volumegroups.NewVolumeGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumeGroups client: %+v", err)
	}
	configureFunc(volumeGroupsClient.Client)

	volumesClient, err := volumes.NewVolumesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Volumes client: %+v", err)
	}
	configureFunc(volumesClient.Client)

	return &Client{
		ElasticSan:     elasticSanClient,
		ElasticSanSkus: elasticSanSkusClient,
		ElasticSans:    elasticSansClient,
		VolumeGroups:   volumeGroupsClient,
		Volumes:        volumesClient,
	}, nil
}
