package v2022_03_02

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-02/diskaccesses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-02/diskencryptionsets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-02/disks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-02/incrementalrestorepoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-02/snapshots"
)

type Client struct {
	DiskAccesses             *diskaccesses.DiskAccessesClient
	DiskEncryptionSets       *diskencryptionsets.DiskEncryptionSetsClient
	Disks                    *disks.DisksClient
	IncrementalRestorePoints *incrementalrestorepoints.IncrementalRestorePointsClient
	Snapshots                *snapshots.SnapshotsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	diskAccessesClient := diskaccesses.NewDiskAccessesClientWithBaseURI(endpoint)
	configureAuthFunc(&diskAccessesClient.Client)

	diskEncryptionSetsClient := diskencryptionsets.NewDiskEncryptionSetsClientWithBaseURI(endpoint)
	configureAuthFunc(&diskEncryptionSetsClient.Client)

	disksClient := disks.NewDisksClientWithBaseURI(endpoint)
	configureAuthFunc(&disksClient.Client)

	incrementalRestorePointsClient := incrementalrestorepoints.NewIncrementalRestorePointsClientWithBaseURI(endpoint)
	configureAuthFunc(&incrementalRestorePointsClient.Client)

	snapshotsClient := snapshots.NewSnapshotsClientWithBaseURI(endpoint)
	configureAuthFunc(&snapshotsClient.Client)

	return Client{
		DiskAccesses:             &diskAccessesClient,
		DiskEncryptionSets:       &diskEncryptionSetsClient,
		Disks:                    &disksClient,
		IncrementalRestorePoints: &incrementalRestorePointsClient,
		Snapshots:                &snapshotsClient,
	}
}
