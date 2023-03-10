package v2021_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagepool/2021-08-01/diskpools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagepool/2021-08-01/diskpoolzones"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagepool/2021-08-01/iscsitargets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagepool/2021-08-01/resourceskus"
)

type Client struct {
	DiskPoolZones *diskpoolzones.DiskPoolZonesClient
	DiskPools     *diskpools.DiskPoolsClient
	IscsiTargets  *iscsitargets.IscsiTargetsClient
	ResourceSkus  *resourceskus.ResourceSkusClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	diskPoolZonesClient := diskpoolzones.NewDiskPoolZonesClientWithBaseURI(endpoint)
	configureAuthFunc(&diskPoolZonesClient.Client)

	diskPoolsClient := diskpools.NewDiskPoolsClientWithBaseURI(endpoint)
	configureAuthFunc(&diskPoolsClient.Client)

	iscsiTargetsClient := iscsitargets.NewIscsiTargetsClientWithBaseURI(endpoint)
	configureAuthFunc(&iscsiTargetsClient.Client)

	resourceSkusClient := resourceskus.NewResourceSkusClientWithBaseURI(endpoint)
	configureAuthFunc(&resourceSkusClient.Client)

	return Client{
		DiskPoolZones: &diskPoolZonesClient,
		DiskPools:     &diskPoolsClient,
		IscsiTargets:  &iscsiTargetsClient,
		ResourceSkus:  &resourceSkusClient,
	}
}
