package v2022_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/backuppolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/capacitypools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/filelocks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/netappaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/netappresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/poolchange"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/resetcifspassword"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/restore"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/snapshotpolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/snapshotpolicylistvolumes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/snapshots"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/subvolumes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/volumegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/volumequotarules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/volumes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/volumesrelocation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/volumesreplication"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/volumesrevert"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	BackupPolicy              *backuppolicy.BackupPolicyClient
	Backups                   *backups.BackupsClient
	CapacityPools             *capacitypools.CapacityPoolsClient
	FileLocks                 *filelocks.FileLocksClient
	NetAppAccounts            *netappaccounts.NetAppAccountsClient
	NetAppResource            *netappresource.NetAppResourceClient
	PoolChange                *poolchange.PoolChangeClient
	ResetCifsPassword         *resetcifspassword.ResetCifsPasswordClient
	Restore                   *restore.RestoreClient
	SnapshotPolicy            *snapshotpolicy.SnapshotPolicyClient
	SnapshotPolicyListVolumes *snapshotpolicylistvolumes.SnapshotPolicyListVolumesClient
	Snapshots                 *snapshots.SnapshotsClient
	SubVolumes                *subvolumes.SubVolumesClient
	VolumeGroups              *volumegroups.VolumeGroupsClient
	VolumeQuotaRules          *volumequotarules.VolumeQuotaRulesClient
	Volumes                   *volumes.VolumesClient
	VolumesRelocation         *volumesrelocation.VolumesRelocationClient
	VolumesReplication        *volumesreplication.VolumesReplicationClient
	VolumesRevert             *volumesrevert.VolumesRevertClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	backupPolicyClient, err := backuppolicy.NewBackupPolicyClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building BackupPolicy client: %+v", err)
	}
	configureFunc(backupPolicyClient.Client)

	backupsClient, err := backups.NewBackupsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Backups client: %+v", err)
	}
	configureFunc(backupsClient.Client)

	capacityPoolsClient, err := capacitypools.NewCapacityPoolsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CapacityPools client: %+v", err)
	}
	configureFunc(capacityPoolsClient.Client)

	fileLocksClient, err := filelocks.NewFileLocksClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building FileLocks client: %+v", err)
	}
	configureFunc(fileLocksClient.Client)

	netAppAccountsClient, err := netappaccounts.NewNetAppAccountsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NetAppAccounts client: %+v", err)
	}
	configureFunc(netAppAccountsClient.Client)

	netAppResourceClient, err := netappresource.NewNetAppResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NetAppResource client: %+v", err)
	}
	configureFunc(netAppResourceClient.Client)

	poolChangeClient, err := poolchange.NewPoolChangeClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PoolChange client: %+v", err)
	}
	configureFunc(poolChangeClient.Client)

	resetCifsPasswordClient, err := resetcifspassword.NewResetCifsPasswordClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ResetCifsPassword client: %+v", err)
	}
	configureFunc(resetCifsPasswordClient.Client)

	restoreClient, err := restore.NewRestoreClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Restore client: %+v", err)
	}
	configureFunc(restoreClient.Client)

	snapshotPolicyClient, err := snapshotpolicy.NewSnapshotPolicyClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SnapshotPolicy client: %+v", err)
	}
	configureFunc(snapshotPolicyClient.Client)

	snapshotPolicyListVolumesClient, err := snapshotpolicylistvolumes.NewSnapshotPolicyListVolumesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SnapshotPolicyListVolumes client: %+v", err)
	}
	configureFunc(snapshotPolicyListVolumesClient.Client)

	snapshotsClient, err := snapshots.NewSnapshotsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Snapshots client: %+v", err)
	}
	configureFunc(snapshotsClient.Client)

	subVolumesClient, err := subvolumes.NewSubVolumesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SubVolumes client: %+v", err)
	}
	configureFunc(subVolumesClient.Client)

	volumeGroupsClient, err := volumegroups.NewVolumeGroupsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VolumeGroups client: %+v", err)
	}
	configureFunc(volumeGroupsClient.Client)

	volumeQuotaRulesClient, err := volumequotarules.NewVolumeQuotaRulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VolumeQuotaRules client: %+v", err)
	}
	configureFunc(volumeQuotaRulesClient.Client)

	volumesClient, err := volumes.NewVolumesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Volumes client: %+v", err)
	}
	configureFunc(volumesClient.Client)

	volumesRelocationClient, err := volumesrelocation.NewVolumesRelocationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VolumesRelocation client: %+v", err)
	}
	configureFunc(volumesRelocationClient.Client)

	volumesReplicationClient, err := volumesreplication.NewVolumesReplicationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VolumesReplication client: %+v", err)
	}
	configureFunc(volumesReplicationClient.Client)

	volumesRevertClient, err := volumesrevert.NewVolumesRevertClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VolumesRevert client: %+v", err)
	}
	configureFunc(volumesRevertClient.Client)

	return &Client{
		BackupPolicy:              backupPolicyClient,
		Backups:                   backupsClient,
		CapacityPools:             capacityPoolsClient,
		FileLocks:                 fileLocksClient,
		NetAppAccounts:            netAppAccountsClient,
		NetAppResource:            netAppResourceClient,
		PoolChange:                poolChangeClient,
		ResetCifsPassword:         resetCifsPasswordClient,
		Restore:                   restoreClient,
		SnapshotPolicy:            snapshotPolicyClient,
		SnapshotPolicyListVolumes: snapshotPolicyListVolumesClient,
		Snapshots:                 snapshotsClient,
		SubVolumes:                subVolumesClient,
		VolumeGroups:              volumeGroupsClient,
		VolumeQuotaRules:          volumeQuotaRulesClient,
		Volumes:                   volumesClient,
		VolumesRelocation:         volumesRelocationClient,
		VolumesReplication:        volumesReplicationClient,
		VolumesRevert:             volumesRevertClient,
	}, nil
}
