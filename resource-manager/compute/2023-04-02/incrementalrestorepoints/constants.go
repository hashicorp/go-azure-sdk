package incrementalrestorepoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessLevel string

const (
	AccessLevelNone  AccessLevel = "None"
	AccessLevelRead  AccessLevel = "Read"
	AccessLevelWrite AccessLevel = "Write"
)

func PossibleValuesForAccessLevel() []string {
	return []string{
		string(AccessLevelNone),
		string(AccessLevelRead),
		string(AccessLevelWrite),
	}
}

func (s *AccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessLevel(input string) (*AccessLevel, error) {
	vals := map[string]AccessLevel{
		"none":  AccessLevelNone,
		"read":  AccessLevelRead,
		"write": AccessLevelWrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessLevel(input)
	return &out, nil
}

type Architecture string

const (
	ArchitectureArmSixFour Architecture = "Arm64"
	ArchitectureXSixFour   Architecture = "x64"
)

func PossibleValuesForArchitecture() []string {
	return []string{
		string(ArchitectureArmSixFour),
		string(ArchitectureXSixFour),
	}
}

func (s *Architecture) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseArchitecture(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseArchitecture(input string) (*Architecture, error) {
	vals := map[string]Architecture{
		"arm64": ArchitectureArmSixFour,
		"x64":   ArchitectureXSixFour,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Architecture(input)
	return &out, nil
}

type DiskSecurityTypes string

const (
	DiskSecurityTypesConfidentialVMDiskEncryptedWithCustomerKey             DiskSecurityTypes = "ConfidentialVM_DiskEncryptedWithCustomerKey"
	DiskSecurityTypesConfidentialVMDiskEncryptedWithPlatformKey             DiskSecurityTypes = "ConfidentialVM_DiskEncryptedWithPlatformKey"
	DiskSecurityTypesConfidentialVMVMGuestStateOnlyEncryptedWithPlatformKey DiskSecurityTypes = "ConfidentialVM_VMGuestStateOnlyEncryptedWithPlatformKey"
	DiskSecurityTypesTrustedLaunch                                          DiskSecurityTypes = "TrustedLaunch"
)

func PossibleValuesForDiskSecurityTypes() []string {
	return []string{
		string(DiskSecurityTypesConfidentialVMDiskEncryptedWithCustomerKey),
		string(DiskSecurityTypesConfidentialVMDiskEncryptedWithPlatformKey),
		string(DiskSecurityTypesConfidentialVMVMGuestStateOnlyEncryptedWithPlatformKey),
		string(DiskSecurityTypesTrustedLaunch),
	}
}

func (s *DiskSecurityTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiskSecurityTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiskSecurityTypes(input string) (*DiskSecurityTypes, error) {
	vals := map[string]DiskSecurityTypes{
		"confidentialvm_diskencryptedwithcustomerkey":             DiskSecurityTypesConfidentialVMDiskEncryptedWithCustomerKey,
		"confidentialvm_diskencryptedwithplatformkey":             DiskSecurityTypesConfidentialVMDiskEncryptedWithPlatformKey,
		"confidentialvm_vmgueststateonlyencryptedwithplatformkey": DiskSecurityTypesConfidentialVMVMGuestStateOnlyEncryptedWithPlatformKey,
		"trustedlaunch": DiskSecurityTypesTrustedLaunch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiskSecurityTypes(input)
	return &out, nil
}

type EncryptionType string

const (
	EncryptionTypeEncryptionAtRestWithCustomerKey             EncryptionType = "EncryptionAtRestWithCustomerKey"
	EncryptionTypeEncryptionAtRestWithPlatformAndCustomerKeys EncryptionType = "EncryptionAtRestWithPlatformAndCustomerKeys"
	EncryptionTypeEncryptionAtRestWithPlatformKey             EncryptionType = "EncryptionAtRestWithPlatformKey"
)

func PossibleValuesForEncryptionType() []string {
	return []string{
		string(EncryptionTypeEncryptionAtRestWithCustomerKey),
		string(EncryptionTypeEncryptionAtRestWithPlatformAndCustomerKeys),
		string(EncryptionTypeEncryptionAtRestWithPlatformKey),
	}
}

func (s *EncryptionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptionType(input string) (*EncryptionType, error) {
	vals := map[string]EncryptionType{
		"encryptionatrestwithcustomerkey":             EncryptionTypeEncryptionAtRestWithCustomerKey,
		"encryptionatrestwithplatformandcustomerkeys": EncryptionTypeEncryptionAtRestWithPlatformAndCustomerKeys,
		"encryptionatrestwithplatformkey":             EncryptionTypeEncryptionAtRestWithPlatformKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionType(input)
	return &out, nil
}

type FileFormat string

const (
	FileFormatVHD  FileFormat = "VHD"
	FileFormatVHDX FileFormat = "VHDX"
)

func PossibleValuesForFileFormat() []string {
	return []string{
		string(FileFormatVHD),
		string(FileFormatVHDX),
	}
}

func (s *FileFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileFormat(input string) (*FileFormat, error) {
	vals := map[string]FileFormat{
		"vhd":  FileFormatVHD,
		"vhdx": FileFormatVHDX,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileFormat(input)
	return &out, nil
}

type HyperVGeneration string

const (
	HyperVGenerationVOne HyperVGeneration = "V1"
	HyperVGenerationVTwo HyperVGeneration = "V2"
)

func PossibleValuesForHyperVGeneration() []string {
	return []string{
		string(HyperVGenerationVOne),
		string(HyperVGenerationVTwo),
	}
}

func (s *HyperVGeneration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHyperVGeneration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHyperVGeneration(input string) (*HyperVGeneration, error) {
	vals := map[string]HyperVGeneration{
		"v1": HyperVGenerationVOne,
		"v2": HyperVGenerationVTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HyperVGeneration(input)
	return &out, nil
}

type NetworkAccessPolicy string

const (
	NetworkAccessPolicyAllowAll     NetworkAccessPolicy = "AllowAll"
	NetworkAccessPolicyAllowPrivate NetworkAccessPolicy = "AllowPrivate"
	NetworkAccessPolicyDenyAll      NetworkAccessPolicy = "DenyAll"
)

func PossibleValuesForNetworkAccessPolicy() []string {
	return []string{
		string(NetworkAccessPolicyAllowAll),
		string(NetworkAccessPolicyAllowPrivate),
		string(NetworkAccessPolicyDenyAll),
	}
}

func (s *NetworkAccessPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkAccessPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkAccessPolicy(input string) (*NetworkAccessPolicy, error) {
	vals := map[string]NetworkAccessPolicy{
		"allowall":     NetworkAccessPolicyAllowAll,
		"allowprivate": NetworkAccessPolicyAllowPrivate,
		"denyall":      NetworkAccessPolicyDenyAll,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkAccessPolicy(input)
	return &out, nil
}

type OperatingSystemTypes string

const (
	OperatingSystemTypesLinux   OperatingSystemTypes = "Linux"
	OperatingSystemTypesWindows OperatingSystemTypes = "Windows"
)

func PossibleValuesForOperatingSystemTypes() []string {
	return []string{
		string(OperatingSystemTypesLinux),
		string(OperatingSystemTypesWindows),
	}
}

func (s *OperatingSystemTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperatingSystemTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperatingSystemTypes(input string) (*OperatingSystemTypes, error) {
	vals := map[string]OperatingSystemTypes{
		"linux":   OperatingSystemTypesLinux,
		"windows": OperatingSystemTypesWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperatingSystemTypes(input)
	return &out, nil
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}

func (s *PublicNetworkAccess) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePublicNetworkAccess(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePublicNetworkAccess(input string) (*PublicNetworkAccess, error) {
	vals := map[string]PublicNetworkAccess{
		"disabled": PublicNetworkAccessDisabled,
		"enabled":  PublicNetworkAccessEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicNetworkAccess(input)
	return &out, nil
}
