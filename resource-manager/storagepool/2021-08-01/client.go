package v2021_08_01

import (
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
