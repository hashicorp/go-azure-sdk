// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_01_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/backuppolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/capacitypools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/netappaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/netappresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/poolchange"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/resetcifspassword"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/restore"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/snapshotpolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/snapshotpolicylistvolumes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/snapshots"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/subvolumes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/vaults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/volumegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/volumequotarules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/volumes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/volumesrelocation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/volumesreplication"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/volumesrevert"
)

type Client struct {
	BackupPolicy              *backuppolicy.BackupPolicyClient
	Backups                   *backups.BackupsClient
	CapacityPools             *capacitypools.CapacityPoolsClient
	NetAppAccounts            *netappaccounts.NetAppAccountsClient
	NetAppResource            *netappresource.NetAppResourceClient
	PoolChange                *poolchange.PoolChangeClient
	ResetCifsPassword         *resetcifspassword.ResetCifsPasswordClient
	Restore                   *restore.RestoreClient
	SnapshotPolicy            *snapshotpolicy.SnapshotPolicyClient
	SnapshotPolicyListVolumes *snapshotpolicylistvolumes.SnapshotPolicyListVolumesClient
	Snapshots                 *snapshots.SnapshotsClient
	SubVolumes                *subvolumes.SubVolumesClient
	Vaults                    *vaults.VaultsClient
	VolumeGroups              *volumegroups.VolumeGroupsClient
	VolumeQuotaRules          *volumequotarules.VolumeQuotaRulesClient
	Volumes                   *volumes.VolumesClient
	VolumesRelocation         *volumesrelocation.VolumesRelocationClient
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

	resetCifsPasswordClient := resetcifspassword.NewResetCifsPasswordClientWithBaseURI(endpoint)
	configureAuthFunc(&resetCifsPasswordClient.Client)

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

	volumeQuotaRulesClient := volumequotarules.NewVolumeQuotaRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&volumeQuotaRulesClient.Client)

	volumesClient := volumes.NewVolumesClientWithBaseURI(endpoint)
	configureAuthFunc(&volumesClient.Client)

	volumesRelocationClient := volumesrelocation.NewVolumesRelocationClientWithBaseURI(endpoint)
	configureAuthFunc(&volumesRelocationClient.Client)

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
		ResetCifsPassword:         &resetCifsPasswordClient,
		Restore:                   &restoreClient,
		SnapshotPolicy:            &snapshotPolicyClient,
		SnapshotPolicyListVolumes: &snapshotPolicyListVolumesClient,
		Snapshots:                 &snapshotsClient,
		SubVolumes:                &subVolumesClient,
		Vaults:                    &vaultsClient,
		VolumeGroups:              &volumeGroupsClient,
		VolumeQuotaRules:          &volumeQuotaRulesClient,
		Volumes:                   &volumesClient,
		VolumesRelocation:         &volumesRelocationClient,
		VolumesReplication:        &volumesReplicationClient,
		VolumesRevert:             &volumesRevertClient,
	}
}
