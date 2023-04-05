package netappresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameResourceTypes string

const (
	CheckNameResourceTypesMicrosoftPointNetAppNetAppAccounts                              CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts"
	CheckNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPools                 CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools"
	CheckNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPoolsVolumes          CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes"
	CheckNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPoolsVolumesSnapshots CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots"
)

func PossibleValuesForCheckNameResourceTypes() []string {
	return []string{
		string(CheckNameResourceTypesMicrosoftPointNetAppNetAppAccounts),
		string(CheckNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPools),
		string(CheckNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPoolsVolumes),
		string(CheckNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPoolsVolumesSnapshots),
	}
}

type CheckQuotaNameResourceTypes string

const (
	CheckQuotaNameResourceTypesMicrosoftPointNetAppNetAppAccounts                              CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts"
	CheckQuotaNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPools                 CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools"
	CheckQuotaNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPoolsVolumes          CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes"
	CheckQuotaNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPoolsVolumesSnapshots CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots"
)

func PossibleValuesForCheckQuotaNameResourceTypes() []string {
	return []string{
		string(CheckQuotaNameResourceTypesMicrosoftPointNetAppNetAppAccounts),
		string(CheckQuotaNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPools),
		string(CheckQuotaNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPoolsVolumes),
		string(CheckQuotaNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPoolsVolumesSnapshots),
	}
}

type InAvailabilityReasonType string

const (
	InAvailabilityReasonTypeAlreadyExists InAvailabilityReasonType = "AlreadyExists"
	InAvailabilityReasonTypeInvalid       InAvailabilityReasonType = "Invalid"
)

func PossibleValuesForInAvailabilityReasonType() []string {
	return []string{
		string(InAvailabilityReasonTypeAlreadyExists),
		string(InAvailabilityReasonTypeInvalid),
	}
}

type RegionStorageToNetworkProximity string

const (
	RegionStorageToNetworkProximityDefault     RegionStorageToNetworkProximity = "Default"
	RegionStorageToNetworkProximityTOne        RegionStorageToNetworkProximity = "T1"
	RegionStorageToNetworkProximityTOneAndTTwo RegionStorageToNetworkProximity = "T1AndT2"
	RegionStorageToNetworkProximityTTwo        RegionStorageToNetworkProximity = "T2"
)

func PossibleValuesForRegionStorageToNetworkProximity() []string {
	return []string{
		string(RegionStorageToNetworkProximityDefault),
		string(RegionStorageToNetworkProximityTOne),
		string(RegionStorageToNetworkProximityTOneAndTTwo),
		string(RegionStorageToNetworkProximityTTwo),
	}
}
