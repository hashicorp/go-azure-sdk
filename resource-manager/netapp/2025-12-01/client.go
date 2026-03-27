package v2025_12_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/backuppolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/backupvaults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/capacitypools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/netappaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/netappresourcequotalimitsaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/netapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/ransomwarereports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/regioninforesources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/snapshotpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/snapshots"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/subscriptionquotaitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/subvolumeinfos"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/volumegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/volumequotarules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/volumes"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	BackupPolicies                   *backuppolicies.BackupPoliciesClient
	BackupVaults                     *backupvaults.BackupVaultsClient
	Backups                          *backups.BackupsClient
	CapacityPools                    *capacitypools.CapacityPoolsClient
	NetAppAccounts                   *netappaccounts.NetAppAccountsClient
	NetAppResourceQuotaLimitsAccount *netappresourcequotalimitsaccount.NetAppResourceQuotaLimitsAccountClient
	Netapps                          *netapps.NetappsClient
	RansomwareReports                *ransomwarereports.RansomwareReportsClient
	RegionInfoResources              *regioninforesources.RegionInfoResourcesClient
	SnapshotPolicies                 *snapshotpolicies.SnapshotPoliciesClient
	Snapshots                        *snapshots.SnapshotsClient
	SubscriptionQuotaItems           *subscriptionquotaitems.SubscriptionQuotaItemsClient
	SubvolumeInfos                   *subvolumeinfos.SubvolumeInfosClient
	VolumeGroups                     *volumegroups.VolumeGroupsClient
	VolumeQuotaRules                 *volumequotarules.VolumeQuotaRulesClient
	Volumes                          *volumes.VolumesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	backupPoliciesClient, err := backuppolicies.NewBackupPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupPolicies client: %+v", err)
	}
	configureFunc(backupPoliciesClient.Client)

	backupVaultsClient, err := backupvaults.NewBackupVaultsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupVaults client: %+v", err)
	}
	configureFunc(backupVaultsClient.Client)

	backupsClient, err := backups.NewBackupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Backups client: %+v", err)
	}
	configureFunc(backupsClient.Client)

	capacityPoolsClient, err := capacitypools.NewCapacityPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapacityPools client: %+v", err)
	}
	configureFunc(capacityPoolsClient.Client)

	netAppAccountsClient, err := netappaccounts.NewNetAppAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetAppAccounts client: %+v", err)
	}
	configureFunc(netAppAccountsClient.Client)

	netAppResourceQuotaLimitsAccountClient, err := netappresourcequotalimitsaccount.NewNetAppResourceQuotaLimitsAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetAppResourceQuotaLimitsAccount client: %+v", err)
	}
	configureFunc(netAppResourceQuotaLimitsAccountClient.Client)

	netappsClient, err := netapps.NewNetappsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Netapps client: %+v", err)
	}
	configureFunc(netappsClient.Client)

	ransomwareReportsClient, err := ransomwarereports.NewRansomwareReportsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RansomwareReports client: %+v", err)
	}
	configureFunc(ransomwareReportsClient.Client)

	regionInfoResourcesClient, err := regioninforesources.NewRegionInfoResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RegionInfoResources client: %+v", err)
	}
	configureFunc(regionInfoResourcesClient.Client)

	snapshotPoliciesClient, err := snapshotpolicies.NewSnapshotPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SnapshotPolicies client: %+v", err)
	}
	configureFunc(snapshotPoliciesClient.Client)

	snapshotsClient, err := snapshots.NewSnapshotsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Snapshots client: %+v", err)
	}
	configureFunc(snapshotsClient.Client)

	subscriptionQuotaItemsClient, err := subscriptionquotaitems.NewSubscriptionQuotaItemsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubscriptionQuotaItems client: %+v", err)
	}
	configureFunc(subscriptionQuotaItemsClient.Client)

	subvolumeInfosClient, err := subvolumeinfos.NewSubvolumeInfosClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubvolumeInfos client: %+v", err)
	}
	configureFunc(subvolumeInfosClient.Client)

	volumeGroupsClient, err := volumegroups.NewVolumeGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumeGroups client: %+v", err)
	}
	configureFunc(volumeGroupsClient.Client)

	volumeQuotaRulesClient, err := volumequotarules.NewVolumeQuotaRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumeQuotaRules client: %+v", err)
	}
	configureFunc(volumeQuotaRulesClient.Client)

	volumesClient, err := volumes.NewVolumesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Volumes client: %+v", err)
	}
	configureFunc(volumesClient.Client)

	return &Client{
		BackupPolicies:                   backupPoliciesClient,
		BackupVaults:                     backupVaultsClient,
		Backups:                          backupsClient,
		CapacityPools:                    capacityPoolsClient,
		NetAppAccounts:                   netAppAccountsClient,
		NetAppResourceQuotaLimitsAccount: netAppResourceQuotaLimitsAccountClient,
		Netapps:                          netappsClient,
		RansomwareReports:                ransomwareReportsClient,
		RegionInfoResources:              regionInfoResourcesClient,
		SnapshotPolicies:                 snapshotPoliciesClient,
		Snapshots:                        snapshotsClient,
		SubscriptionQuotaItems:           subscriptionQuotaItemsClient,
		SubvolumeInfos:                   subvolumeInfosClient,
		VolumeGroups:                     volumeGroupsClient,
		VolumeQuotaRules:                 volumeQuotaRulesClient,
		Volumes:                          volumesClient,
	}, nil
}
