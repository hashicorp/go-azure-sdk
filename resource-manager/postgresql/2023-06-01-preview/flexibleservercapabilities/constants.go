package flexibleservercapabilities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityStatus string

const (
	CapabilityStatusAvailable CapabilityStatus = "Available"
	CapabilityStatusDefault   CapabilityStatus = "Default"
	CapabilityStatusDisabled  CapabilityStatus = "Disabled"
	CapabilityStatusVisible   CapabilityStatus = "Visible"
)

func PossibleValuesForCapabilityStatus() []string {
	return []string{
		string(CapabilityStatusAvailable),
		string(CapabilityStatusDefault),
		string(CapabilityStatusDisabled),
		string(CapabilityStatusVisible),
	}
}

func (s *CapabilityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCapabilityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCapabilityStatus(input string) (*CapabilityStatus, error) {
	vals := map[string]CapabilityStatus{
		"available": CapabilityStatusAvailable,
		"default":   CapabilityStatusDefault,
		"disabled":  CapabilityStatusDisabled,
		"visible":   CapabilityStatusVisible,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CapabilityStatus(input)
	return &out, nil
}

type FastProvisioningSupportedEnum string

const (
	FastProvisioningSupportedEnumDisabled FastProvisioningSupportedEnum = "Disabled"
	FastProvisioningSupportedEnumEnabled  FastProvisioningSupportedEnum = "Enabled"
)

func PossibleValuesForFastProvisioningSupportedEnum() []string {
	return []string{
		string(FastProvisioningSupportedEnumDisabled),
		string(FastProvisioningSupportedEnumEnabled),
	}
}

func (s *FastProvisioningSupportedEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFastProvisioningSupportedEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFastProvisioningSupportedEnum(input string) (*FastProvisioningSupportedEnum, error) {
	vals := map[string]FastProvisioningSupportedEnum{
		"disabled": FastProvisioningSupportedEnumDisabled,
		"enabled":  FastProvisioningSupportedEnumEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FastProvisioningSupportedEnum(input)
	return &out, nil
}

type GeoBackupSupportedEnum string

const (
	GeoBackupSupportedEnumDisabled GeoBackupSupportedEnum = "Disabled"
	GeoBackupSupportedEnumEnabled  GeoBackupSupportedEnum = "Enabled"
)

func PossibleValuesForGeoBackupSupportedEnum() []string {
	return []string{
		string(GeoBackupSupportedEnumDisabled),
		string(GeoBackupSupportedEnumEnabled),
	}
}

func (s *GeoBackupSupportedEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGeoBackupSupportedEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGeoBackupSupportedEnum(input string) (*GeoBackupSupportedEnum, error) {
	vals := map[string]GeoBackupSupportedEnum{
		"disabled": GeoBackupSupportedEnumDisabled,
		"enabled":  GeoBackupSupportedEnumEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GeoBackupSupportedEnum(input)
	return &out, nil
}

type HaMode string

const (
	HaModeSameZone      HaMode = "SameZone"
	HaModeZoneRedundant HaMode = "ZoneRedundant"
)

func PossibleValuesForHaMode() []string {
	return []string{
		string(HaModeSameZone),
		string(HaModeZoneRedundant),
	}
}

func (s *HaMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHaMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHaMode(input string) (*HaMode, error) {
	vals := map[string]HaMode{
		"samezone":      HaModeSameZone,
		"zoneredundant": HaModeZoneRedundant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HaMode(input)
	return &out, nil
}

type OnlineResizeSupportedEnum string

const (
	OnlineResizeSupportedEnumDisabled OnlineResizeSupportedEnum = "Disabled"
	OnlineResizeSupportedEnumEnabled  OnlineResizeSupportedEnum = "Enabled"
)

func PossibleValuesForOnlineResizeSupportedEnum() []string {
	return []string{
		string(OnlineResizeSupportedEnumDisabled),
		string(OnlineResizeSupportedEnumEnabled),
	}
}

func (s *OnlineResizeSupportedEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineResizeSupportedEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineResizeSupportedEnum(input string) (*OnlineResizeSupportedEnum, error) {
	vals := map[string]OnlineResizeSupportedEnum{
		"disabled": OnlineResizeSupportedEnumDisabled,
		"enabled":  OnlineResizeSupportedEnumEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineResizeSupportedEnum(input)
	return &out, nil
}

type RestrictedEnum string

const (
	RestrictedEnumDisabled RestrictedEnum = "Disabled"
	RestrictedEnumEnabled  RestrictedEnum = "Enabled"
)

func PossibleValuesForRestrictedEnum() []string {
	return []string{
		string(RestrictedEnumDisabled),
		string(RestrictedEnumEnabled),
	}
}

func (s *RestrictedEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestrictedEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestrictedEnum(input string) (*RestrictedEnum, error) {
	vals := map[string]RestrictedEnum{
		"disabled": RestrictedEnumDisabled,
		"enabled":  RestrictedEnumEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestrictedEnum(input)
	return &out, nil
}

type StorageAutoGrowthSupportedEnum string

const (
	StorageAutoGrowthSupportedEnumDisabled StorageAutoGrowthSupportedEnum = "Disabled"
	StorageAutoGrowthSupportedEnumEnabled  StorageAutoGrowthSupportedEnum = "Enabled"
)

func PossibleValuesForStorageAutoGrowthSupportedEnum() []string {
	return []string{
		string(StorageAutoGrowthSupportedEnumDisabled),
		string(StorageAutoGrowthSupportedEnumEnabled),
	}
}

func (s *StorageAutoGrowthSupportedEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageAutoGrowthSupportedEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageAutoGrowthSupportedEnum(input string) (*StorageAutoGrowthSupportedEnum, error) {
	vals := map[string]StorageAutoGrowthSupportedEnum{
		"disabled": StorageAutoGrowthSupportedEnumDisabled,
		"enabled":  StorageAutoGrowthSupportedEnumEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageAutoGrowthSupportedEnum(input)
	return &out, nil
}

type ZoneRedundantHaAndGeoBackupSupportedEnum string

const (
	ZoneRedundantHaAndGeoBackupSupportedEnumDisabled ZoneRedundantHaAndGeoBackupSupportedEnum = "Disabled"
	ZoneRedundantHaAndGeoBackupSupportedEnumEnabled  ZoneRedundantHaAndGeoBackupSupportedEnum = "Enabled"
)

func PossibleValuesForZoneRedundantHaAndGeoBackupSupportedEnum() []string {
	return []string{
		string(ZoneRedundantHaAndGeoBackupSupportedEnumDisabled),
		string(ZoneRedundantHaAndGeoBackupSupportedEnumEnabled),
	}
}

func (s *ZoneRedundantHaAndGeoBackupSupportedEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZoneRedundantHaAndGeoBackupSupportedEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZoneRedundantHaAndGeoBackupSupportedEnum(input string) (*ZoneRedundantHaAndGeoBackupSupportedEnum, error) {
	vals := map[string]ZoneRedundantHaAndGeoBackupSupportedEnum{
		"disabled": ZoneRedundantHaAndGeoBackupSupportedEnumDisabled,
		"enabled":  ZoneRedundantHaAndGeoBackupSupportedEnumEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZoneRedundantHaAndGeoBackupSupportedEnum(input)
	return &out, nil
}

type ZoneRedundantHaSupportedEnum string

const (
	ZoneRedundantHaSupportedEnumDisabled ZoneRedundantHaSupportedEnum = "Disabled"
	ZoneRedundantHaSupportedEnumEnabled  ZoneRedundantHaSupportedEnum = "Enabled"
)

func PossibleValuesForZoneRedundantHaSupportedEnum() []string {
	return []string{
		string(ZoneRedundantHaSupportedEnumDisabled),
		string(ZoneRedundantHaSupportedEnumEnabled),
	}
}

func (s *ZoneRedundantHaSupportedEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZoneRedundantHaSupportedEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZoneRedundantHaSupportedEnum(input string) (*ZoneRedundantHaSupportedEnum, error) {
	vals := map[string]ZoneRedundantHaSupportedEnum{
		"disabled": ZoneRedundantHaSupportedEnumDisabled,
		"enabled":  ZoneRedundantHaSupportedEnumEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZoneRedundantHaSupportedEnum(input)
	return &out, nil
}
