package netappresource

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *CheckNameResourceTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCheckNameResourceTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCheckNameResourceTypes(input string) (*CheckNameResourceTypes, error) {
	vals := map[string]CheckNameResourceTypes{
		"microsoft.netapp/netappaccounts":                                 CheckNameResourceTypesMicrosoftPointNetAppNetAppAccounts,
		"microsoft.netapp/netappaccounts/capacitypools":                   CheckNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPools,
		"microsoft.netapp/netappaccounts/capacitypools/volumes":           CheckNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPoolsVolumes,
		"microsoft.netapp/netappaccounts/capacitypools/volumes/snapshots": CheckNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPoolsVolumesSnapshots,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CheckNameResourceTypes(input)
	return &out, nil
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

func (s *CheckQuotaNameResourceTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCheckQuotaNameResourceTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCheckQuotaNameResourceTypes(input string) (*CheckQuotaNameResourceTypes, error) {
	vals := map[string]CheckQuotaNameResourceTypes{
		"microsoft.netapp/netappaccounts":                                 CheckQuotaNameResourceTypesMicrosoftPointNetAppNetAppAccounts,
		"microsoft.netapp/netappaccounts/capacitypools":                   CheckQuotaNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPools,
		"microsoft.netapp/netappaccounts/capacitypools/volumes":           CheckQuotaNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPoolsVolumes,
		"microsoft.netapp/netappaccounts/capacitypools/volumes/snapshots": CheckQuotaNameResourceTypesMicrosoftPointNetAppNetAppAccountsCapacityPoolsVolumesSnapshots,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CheckQuotaNameResourceTypes(input)
	return &out, nil
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

func (s *InAvailabilityReasonType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInAvailabilityReasonType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInAvailabilityReasonType(input string) (*InAvailabilityReasonType, error) {
	vals := map[string]InAvailabilityReasonType{
		"alreadyexists": InAvailabilityReasonTypeAlreadyExists,
		"invalid":       InAvailabilityReasonTypeInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InAvailabilityReasonType(input)
	return &out, nil
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

func (s *RegionStorageToNetworkProximity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegionStorageToNetworkProximity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegionStorageToNetworkProximity(input string) (*RegionStorageToNetworkProximity, error) {
	vals := map[string]RegionStorageToNetworkProximity{
		"default": RegionStorageToNetworkProximityDefault,
		"t1":      RegionStorageToNetworkProximityTOne,
		"t1andt2": RegionStorageToNetworkProximityTOneAndTTwo,
		"t2":      RegionStorageToNetworkProximityTTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegionStorageToNetworkProximity(input)
	return &out, nil
}
