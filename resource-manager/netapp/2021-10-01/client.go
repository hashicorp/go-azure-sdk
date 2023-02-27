// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/backuppolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/capacitypools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/netappaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/netappresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/poolchange"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/restore"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/snapshotpolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/snapshotpolicylistvolumes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/snapshots"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/subvolumes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/vaults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/volumegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/volumes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/volumesreplication"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2021-10-01/volumesrevert"
)

type Client struct {
	BackupPolicy              *backuppolicy.BackupPolicyClient
	Backups                   *backups.BackupsClient
	CapacityPools             *capacitypools.CapacityPoolsClient
	NetAppAccounts            *netappaccounts.NetAppAccountsClient
	NetAppResource            *netappresource.NetAppResourceClient
	PoolChange                *poolchange.PoolChangeClient
	Restore                   *restore.RestoreClient
	SnapshotPolicy            *snapshotpolicy.SnapshotPolicyClient
	SnapshotPolicyListVolumes *snapshotpolicylistvolumes.SnapshotPolicyListVolumesClient
	Snapshots                 *snapshots.SnapshotsClient
	SubVolumes                *subvolumes.SubVolumesClient
	Vaults                    *vaults.VaultsClient
	VolumeGroups              *volumegroups.VolumeGroupsClient
	Volumes                   *volumes.VolumesClient
	VolumesReplication        *volumesreplication.VolumesReplicationClient
	VolumesRevert             *volumesrevert.VolumesRevertClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	backupPolicyClient := backuppolicy.NewBackupPolicyClientWithBaseURI(endpoint)
	configureAuthFunc(&backupPolicyClient.Client)

	backupsClient := backups.NewBackupsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupsClient.Client)

	capacityPoolsClient := capacitypools.NewCapacityPoolsClientWithBaseURI(endpoint)
	configureAuthFunc(&capacityPoolsClient.Client)

	netAppAccountsClient := netappaccounts.NewNetAppAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&netAppAccountsClient.Client)

	netAppResourceClient := netappresource.NewNetAppResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&netAppResourceClient.Client)

	poolChangeClient := poolchange.NewPoolChangeClientWithBaseURI(endpoint)
	configureAuthFunc(&poolChangeClient.Client)

	restoreClient := restore.NewRestoreClientWithBaseURI(endpoint)
	configureAuthFunc(&restoreClient.Client)

	snapshotPolicyClient := snapshotpolicy.NewSnapshotPolicyClientWithBaseURI(endpoint)
	configureAuthFunc(&snapshotPolicyClient.Client)

	snapshotPolicyListVolumesClient := snapshotpolicylistvolumes.NewSnapshotPolicyListVolumesClientWithBaseURI(endpoint)
	configureAuthFunc(&snapshotPolicyListVolumesClient.Client)

	snapshotsClient := snapshots.NewSnapshotsClientWithBaseURI(endpoint)
	configureAuthFunc(&snapshotsClient.Client)

	subVolumesClient := subvolumes.NewSubVolumesClientWithBaseURI(endpoint)
	configureAuthFunc(&subVolumesClient.Client)

	vaultsClient := vaults.NewVaultsClientWithBaseURI(endpoint)
	configureAuthFunc(&vaultsClient.Client)

	volumeGroupsClient := volumegroups.NewVolumeGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&volumeGroupsClient.Client)

	volumesClient := volumes.NewVolumesClientWithBaseURI(endpoint)
	configureAuthFunc(&volumesClient.Client)

	volumesReplicationClient := volumesreplication.NewVolumesReplicationClientWithBaseURI(endpoint)
	configureAuthFunc(&volumesReplicationClient.Client)

	volumesRevertClient := volumesrevert.NewVolumesRevertClientWithBaseURI(endpoint)
	configureAuthFunc(&volumesRevertClient.Client)

	return Client{
		BackupPolicy:              &backupPolicyClient,
		Backups:                   &backupsClient,
		CapacityPools:             &capacityPoolsClient,
		NetAppAccounts:            &netAppAccountsClient,
		NetAppResource:            &netAppResourceClient,
		PoolChange:                &poolChangeClient,
		Restore:                   &restoreClient,
		SnapshotPolicy:            &snapshotPolicyClient,
		SnapshotPolicyListVolumes: &snapshotPolicyListVolumesClient,
		Snapshots:                 &snapshotsClient,
		SubVolumes:                &subVolumesClient,
		Vaults:                    &vaultsClient,
		VolumeGroups:              &volumeGroupsClient,
		Volumes:                   &volumesClient,
		VolumesReplication:        &volumesReplicationClient,
		VolumesRevert:             &volumesRevertClient,
	}
}
