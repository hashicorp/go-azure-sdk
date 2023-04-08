package clusters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaysOfWeek string

const (
	DaysOfWeekFriday    DaysOfWeek = "Friday"
	DaysOfWeekMonday    DaysOfWeek = "Monday"
	DaysOfWeekSaturday  DaysOfWeek = "Saturday"
	DaysOfWeekSunday    DaysOfWeek = "Sunday"
	DaysOfWeekThursday  DaysOfWeek = "Thursday"
	DaysOfWeekTuesday   DaysOfWeek = "Tuesday"
	DaysOfWeekWednesday DaysOfWeek = "Wednesday"
)

func PossibleValuesForDaysOfWeek() []string {
	return []string{
		string(DaysOfWeekFriday),
		string(DaysOfWeekMonday),
		string(DaysOfWeekSaturday),
		string(DaysOfWeekSunday),
		string(DaysOfWeekThursday),
		string(DaysOfWeekTuesday),
		string(DaysOfWeekWednesday),
	}
}

func (s *DaysOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDaysOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDaysOfWeek(input string) (*DaysOfWeek, error) {
	vals := map[string]DaysOfWeek{
		"friday":    DaysOfWeekFriday,
		"monday":    DaysOfWeekMonday,
		"saturday":  DaysOfWeekSaturday,
		"sunday":    DaysOfWeekSunday,
		"thursday":  DaysOfWeekThursday,
		"tuesday":   DaysOfWeekTuesday,
		"wednesday": DaysOfWeekWednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DaysOfWeek(input)
	return &out, nil
}

type DirectoryType string

const (
	DirectoryTypeActiveDirectory DirectoryType = "ActiveDirectory"
)

func PossibleValuesForDirectoryType() []string {
	return []string{
		string(DirectoryTypeActiveDirectory),
	}
}

func (s *DirectoryType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDirectoryType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDirectoryType(input string) (*DirectoryType, error) {
	vals := map[string]DirectoryType{
		"activedirectory": DirectoryTypeActiveDirectory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DirectoryType(input)
	return &out, nil
}

type HDInsightClusterProvisioningState string

const (
	HDInsightClusterProvisioningStateCanceled   HDInsightClusterProvisioningState = "Canceled"
	HDInsightClusterProvisioningStateDeleting   HDInsightClusterProvisioningState = "Deleting"
	HDInsightClusterProvisioningStateFailed     HDInsightClusterProvisioningState = "Failed"
	HDInsightClusterProvisioningStateInProgress HDInsightClusterProvisioningState = "InProgress"
	HDInsightClusterProvisioningStateSucceeded  HDInsightClusterProvisioningState = "Succeeded"
)

func PossibleValuesForHDInsightClusterProvisioningState() []string {
	return []string{
		string(HDInsightClusterProvisioningStateCanceled),
		string(HDInsightClusterProvisioningStateDeleting),
		string(HDInsightClusterProvisioningStateFailed),
		string(HDInsightClusterProvisioningStateInProgress),
		string(HDInsightClusterProvisioningStateSucceeded),
	}
}

func (s *HDInsightClusterProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHDInsightClusterProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHDInsightClusterProvisioningState(input string) (*HDInsightClusterProvisioningState, error) {
	vals := map[string]HDInsightClusterProvisioningState{
		"canceled":   HDInsightClusterProvisioningStateCanceled,
		"deleting":   HDInsightClusterProvisioningStateDeleting,
		"failed":     HDInsightClusterProvisioningStateFailed,
		"inprogress": HDInsightClusterProvisioningStateInProgress,
		"succeeded":  HDInsightClusterProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HDInsightClusterProvisioningState(input)
	return &out, nil
}

type JsonWebKeyEncryptionAlgorithm string

const (
	JsonWebKeyEncryptionAlgorithmRSANegativeOAEP                   JsonWebKeyEncryptionAlgorithm = "RSA-OAEP"
	JsonWebKeyEncryptionAlgorithmRSANegativeOAEPNegativeTwoFiveSix JsonWebKeyEncryptionAlgorithm = "RSA-OAEP-256"
	JsonWebKeyEncryptionAlgorithmRSAOneFive                        JsonWebKeyEncryptionAlgorithm = "RSA1_5"
)

func PossibleValuesForJsonWebKeyEncryptionAlgorithm() []string {
	return []string{
		string(JsonWebKeyEncryptionAlgorithmRSANegativeOAEP),
		string(JsonWebKeyEncryptionAlgorithmRSANegativeOAEPNegativeTwoFiveSix),
		string(JsonWebKeyEncryptionAlgorithmRSAOneFive),
	}
}

func (s *JsonWebKeyEncryptionAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJsonWebKeyEncryptionAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJsonWebKeyEncryptionAlgorithm(input string) (*JsonWebKeyEncryptionAlgorithm, error) {
	vals := map[string]JsonWebKeyEncryptionAlgorithm{
		"rsa-oaep":     JsonWebKeyEncryptionAlgorithmRSANegativeOAEP,
		"rsa-oaep-256": JsonWebKeyEncryptionAlgorithmRSANegativeOAEPNegativeTwoFiveSix,
		"rsa1_5":       JsonWebKeyEncryptionAlgorithmRSAOneFive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JsonWebKeyEncryptionAlgorithm(input)
	return &out, nil
}

type OSType string

const (
	OSTypeLinux   OSType = "Linux"
	OSTypeWindows OSType = "Windows"
)

func PossibleValuesForOSType() []string {
	return []string{
		string(OSTypeLinux),
		string(OSTypeWindows),
	}
}

func (s *OSType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOSType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOSType(input string) (*OSType, error) {
	vals := map[string]OSType{
		"linux":   OSTypeLinux,
		"windows": OSTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OSType(input)
	return &out, nil
}

type PrivateLink string

const (
	PrivateLinkDisabled PrivateLink = "Disabled"
	PrivateLinkEnabled  PrivateLink = "Enabled"
)

func PossibleValuesForPrivateLink() []string {
	return []string{
		string(PrivateLinkDisabled),
		string(PrivateLinkEnabled),
	}
}

func (s *PrivateLink) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivateLink(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivateLink(input string) (*PrivateLink, error) {
	vals := map[string]PrivateLink{
		"disabled": PrivateLinkDisabled,
		"enabled":  PrivateLinkEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateLink(input)
	return &out, nil
}

type ResourceProviderConnection string

const (
	ResourceProviderConnectionInbound  ResourceProviderConnection = "Inbound"
	ResourceProviderConnectionOutbound ResourceProviderConnection = "Outbound"
)

func PossibleValuesForResourceProviderConnection() []string {
	return []string{
		string(ResourceProviderConnectionInbound),
		string(ResourceProviderConnectionOutbound),
	}
}

func (s *ResourceProviderConnection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceProviderConnection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceProviderConnection(input string) (*ResourceProviderConnection, error) {
	vals := map[string]ResourceProviderConnection{
		"inbound":  ResourceProviderConnectionInbound,
		"outbound": ResourceProviderConnectionOutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceProviderConnection(input)
	return &out, nil
}

type Tier string

const (
	TierPremium  Tier = "Premium"
	TierStandard Tier = "Standard"
)

func PossibleValuesForTier() []string {
	return []string{
		string(TierPremium),
		string(TierStandard),
	}
}

func (s *Tier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTier(input string) (*Tier, error) {
	vals := map[string]Tier{
		"premium":  TierPremium,
		"standard": TierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Tier(input)
	return &out, nil
}
