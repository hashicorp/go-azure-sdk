package v2021_11_20_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elasticsan/2021-11-20-preview/elasticsan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elasticsan/2021-11-20-preview/elasticsans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elasticsan/2021-11-20-preview/elasticsanskus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elasticsan/2021-11-20-preview/volumegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elasticsan/2021-11-20-preview/volumes"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	ElasticSan     *elasticsan.ElasticSanClient
	ElasticSanSkus *elasticsanskus.ElasticSanSkusClient
	ElasticSans    *elasticsans.ElasticSansClient
	VolumeGroups   *volumegroups.VolumeGroupsClient
	Volumes        *volumes.VolumesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	elasticSanClient := elasticsan.NewElasticSanClientWithBaseURI(endpoint)
	configureAuthFunc(&elasticSanClient.Client)

	elasticSanSkusClient := elasticsanskus.NewElasticSanSkusClientWithBaseURI(endpoint)
	configureAuthFunc(&elasticSanSkusClient.Client)

	elasticSansClient := elasticsans.NewElasticSansClientWithBaseURI(endpoint)
	configureAuthFunc(&elasticSansClient.Client)

	volumeGroupsClient := volumegroups.NewVolumeGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&volumeGroupsClient.Client)

	volumesClient := volumes.NewVolumesClientWithBaseURI(endpoint)
	configureAuthFunc(&volumesClient.Client)

	return Client{
		ElasticSan:     &elasticSanClient,
		ElasticSanSkus: &elasticSanSkusClient,
		ElasticSans:    &elasticSansClient,
		VolumeGroups:   &volumeGroupsClient,
		Volumes:        &volumesClient,
	}
}
