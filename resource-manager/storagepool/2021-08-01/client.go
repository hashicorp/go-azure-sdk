package v2021_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/storagepool/2021-08-01/diskpools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagepool/2021-08-01/diskpoolzones"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagepool/2021-08-01/iscsitargets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagepool/2021-08-01/resourceskus"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DiskPoolZones *diskpoolzones.DiskPoolZonesClient
	DiskPools     *diskpools.DiskPoolsClient
	IscsiTargets  *iscsitargets.IscsiTargetsClient
	ResourceSkus  *resourceskus.ResourceSkusClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	diskPoolZonesClient, err := diskpoolzones.NewDiskPoolZonesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiskPoolZones client: %+v", err)
	}
	configureFunc(diskPoolZonesClient.Client)

	diskPoolsClient, err := diskpools.NewDiskPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiskPools client: %+v", err)
	}
	configureFunc(diskPoolsClient.Client)

	iscsiTargetsClient, err := iscsitargets.NewIscsiTargetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IscsiTargets client: %+v", err)
	}
	configureFunc(iscsiTargetsClient.Client)

	resourceSkusClient, err := resourceskus.NewResourceSkusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceSkus client: %+v", err)
	}
	configureFunc(resourceSkusClient.Client)

	return &Client{
		DiskPoolZones: diskPoolZonesClient,
		DiskPools:     diskPoolsClient,
		IscsiTargets:  iscsiTargetsClient,
		ResourceSkus:  resourceSkusClient,
	}, nil
}
