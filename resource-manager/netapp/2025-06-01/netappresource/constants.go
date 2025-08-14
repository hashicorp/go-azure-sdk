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

type NetworkFeatures string

const (
	NetworkFeaturesBasic         NetworkFeatures = "Basic"
	NetworkFeaturesBasicStandard NetworkFeatures = "Basic_Standard"
	NetworkFeaturesStandard      NetworkFeatures = "Standard"
	NetworkFeaturesStandardBasic NetworkFeatures = "Standard_Basic"
)

func PossibleValuesForNetworkFeatures() []string {
	return []string{
		string(NetworkFeaturesBasic),
		string(NetworkFeaturesBasicStandard),
		string(NetworkFeaturesStandard),
		string(NetworkFeaturesStandardBasic),
	}
}

func (s *NetworkFeatures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkFeatures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkFeatures(input string) (*NetworkFeatures, error) {
	vals := map[string]NetworkFeatures{
		"basic":          NetworkFeaturesBasic,
		"basic_standard": NetworkFeaturesBasicStandard,
		"standard":       NetworkFeaturesStandard,
		"standard_basic": NetworkFeaturesStandardBasic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkFeatures(input)
	return &out, nil
}

type NetworkSiblingSetProvisioningState string

const (
	NetworkSiblingSetProvisioningStateCanceled  NetworkSiblingSetProvisioningState = "Canceled"
	NetworkSiblingSetProvisioningStateFailed    NetworkSiblingSetProvisioningState = "Failed"
	NetworkSiblingSetProvisioningStateSucceeded NetworkSiblingSetProvisioningState = "Succeeded"
	NetworkSiblingSetProvisioningStateUpdating  NetworkSiblingSetProvisioningState = "Updating"
)

func PossibleValuesForNetworkSiblingSetProvisioningState() []string {
	return []string{
		string(NetworkSiblingSetProvisioningStateCanceled),
		string(NetworkSiblingSetProvisioningStateFailed),
		string(NetworkSiblingSetProvisioningStateSucceeded),
		string(NetworkSiblingSetProvisioningStateUpdating),
	}
}

func (s *NetworkSiblingSetProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkSiblingSetProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkSiblingSetProvisioningState(input string) (*NetworkSiblingSetProvisioningState, error) {
	vals := map[string]NetworkSiblingSetProvisioningState{
		"canceled":  NetworkSiblingSetProvisioningStateCanceled,
		"failed":    NetworkSiblingSetProvisioningStateFailed,
		"succeeded": NetworkSiblingSetProvisioningStateSucceeded,
		"updating":  NetworkSiblingSetProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkSiblingSetProvisioningState(input)
	return &out, nil
}

type RegionStorageToNetworkProximity string

const (
	RegionStorageToNetworkProximityAcrossTTwo               RegionStorageToNetworkProximity = "AcrossT2"
	RegionStorageToNetworkProximityDefault                  RegionStorageToNetworkProximity = "Default"
	RegionStorageToNetworkProximityTOne                     RegionStorageToNetworkProximity = "T1"
	RegionStorageToNetworkProximityTOneAndAcrossTTwo        RegionStorageToNetworkProximity = "T1AndAcrossT2"
	RegionStorageToNetworkProximityTOneAndTTwo              RegionStorageToNetworkProximity = "T1AndT2"
	RegionStorageToNetworkProximityTOneAndTTwoAndAcrossTTwo RegionStorageToNetworkProximity = "T1AndT2AndAcrossT2"
	RegionStorageToNetworkProximityTTwo                     RegionStorageToNetworkProximity = "T2"
	RegionStorageToNetworkProximityTTwoAndAcrossTTwo        RegionStorageToNetworkProximity = "T2AndAcrossT2"
)

func PossibleValuesForRegionStorageToNetworkProximity() []string {
	return []string{
		string(RegionStorageToNetworkProximityAcrossTTwo),
		string(RegionStorageToNetworkProximityDefault),
		string(RegionStorageToNetworkProximityTOne),
		string(RegionStorageToNetworkProximityTOneAndAcrossTTwo),
		string(RegionStorageToNetworkProximityTOneAndTTwo),
		string(RegionStorageToNetworkProximityTOneAndTTwoAndAcrossTTwo),
		string(RegionStorageToNetworkProximityTTwo),
		string(RegionStorageToNetworkProximityTTwoAndAcrossTTwo),
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
		"acrosst2":           RegionStorageToNetworkProximityAcrossTTwo,
		"default":            RegionStorageToNetworkProximityDefault,
		"t1":                 RegionStorageToNetworkProximityTOne,
		"t1andacrosst2":      RegionStorageToNetworkProximityTOneAndAcrossTTwo,
		"t1andt2":            RegionStorageToNetworkProximityTOneAndTTwo,
		"t1andt2andacrosst2": RegionStorageToNetworkProximityTOneAndTTwoAndAcrossTTwo,
		"t2":                 RegionStorageToNetworkProximityTTwo,
		"t2andacrosst2":      RegionStorageToNetworkProximityTTwoAndAcrossTTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegionStorageToNetworkProximity(input)
	return &out, nil
}
