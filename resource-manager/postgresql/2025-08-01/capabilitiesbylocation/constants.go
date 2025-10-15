package capabilitiesbylocation

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

type FastProvisioningSupport string

const (
	FastProvisioningSupportDisabled FastProvisioningSupport = "Disabled"
	FastProvisioningSupportEnabled  FastProvisioningSupport = "Enabled"
)

func PossibleValuesForFastProvisioningSupport() []string {
	return []string{
		string(FastProvisioningSupportDisabled),
		string(FastProvisioningSupportEnabled),
	}
}

func (s *FastProvisioningSupport) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFastProvisioningSupport(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFastProvisioningSupport(input string) (*FastProvisioningSupport, error) {
	vals := map[string]FastProvisioningSupport{
		"disabled": FastProvisioningSupportDisabled,
		"enabled":  FastProvisioningSupportEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FastProvisioningSupport(input)
	return &out, nil
}

type FeatureStatus string

const (
	FeatureStatusDisabled FeatureStatus = "Disabled"
	FeatureStatusEnabled  FeatureStatus = "Enabled"
)

func PossibleValuesForFeatureStatus() []string {
	return []string{
		string(FeatureStatusDisabled),
		string(FeatureStatusEnabled),
	}
}

func (s *FeatureStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureStatus(input string) (*FeatureStatus, error) {
	vals := map[string]FeatureStatus{
		"disabled": FeatureStatusDisabled,
		"enabled":  FeatureStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureStatus(input)
	return &out, nil
}

type GeographicallyRedundantBackupSupport string

const (
	GeographicallyRedundantBackupSupportDisabled GeographicallyRedundantBackupSupport = "Disabled"
	GeographicallyRedundantBackupSupportEnabled  GeographicallyRedundantBackupSupport = "Enabled"
)

func PossibleValuesForGeographicallyRedundantBackupSupport() []string {
	return []string{
		string(GeographicallyRedundantBackupSupportDisabled),
		string(GeographicallyRedundantBackupSupportEnabled),
	}
}

func (s *GeographicallyRedundantBackupSupport) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGeographicallyRedundantBackupSupport(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGeographicallyRedundantBackupSupport(input string) (*GeographicallyRedundantBackupSupport, error) {
	vals := map[string]GeographicallyRedundantBackupSupport{
		"disabled": GeographicallyRedundantBackupSupportDisabled,
		"enabled":  GeographicallyRedundantBackupSupportEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GeographicallyRedundantBackupSupport(input)
	return &out, nil
}

type HighAvailabilityMode string

const (
	HighAvailabilityModeSameZone      HighAvailabilityMode = "SameZone"
	HighAvailabilityModeZoneRedundant HighAvailabilityMode = "ZoneRedundant"
)

func PossibleValuesForHighAvailabilityMode() []string {
	return []string{
		string(HighAvailabilityModeSameZone),
		string(HighAvailabilityModeZoneRedundant),
	}
}

func (s *HighAvailabilityMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHighAvailabilityMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHighAvailabilityMode(input string) (*HighAvailabilityMode, error) {
	vals := map[string]HighAvailabilityMode{
		"samezone":      HighAvailabilityModeSameZone,
		"zoneredundant": HighAvailabilityModeZoneRedundant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HighAvailabilityMode(input)
	return &out, nil
}

type LocationRestricted string

const (
	LocationRestrictedDisabled LocationRestricted = "Disabled"
	LocationRestrictedEnabled  LocationRestricted = "Enabled"
)

func PossibleValuesForLocationRestricted() []string {
	return []string{
		string(LocationRestrictedDisabled),
		string(LocationRestrictedEnabled),
	}
}

func (s *LocationRestricted) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocationRestricted(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocationRestricted(input string) (*LocationRestricted, error) {
	vals := map[string]LocationRestricted{
		"disabled": LocationRestrictedDisabled,
		"enabled":  LocationRestrictedEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationRestricted(input)
	return &out, nil
}

type OnlineStorageResizeSupport string

const (
	OnlineStorageResizeSupportDisabled OnlineStorageResizeSupport = "Disabled"
	OnlineStorageResizeSupportEnabled  OnlineStorageResizeSupport = "Enabled"
)

func PossibleValuesForOnlineStorageResizeSupport() []string {
	return []string{
		string(OnlineStorageResizeSupportDisabled),
		string(OnlineStorageResizeSupportEnabled),
	}
}

func (s *OnlineStorageResizeSupport) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineStorageResizeSupport(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineStorageResizeSupport(input string) (*OnlineStorageResizeSupport, error) {
	vals := map[string]OnlineStorageResizeSupport{
		"disabled": OnlineStorageResizeSupportDisabled,
		"enabled":  OnlineStorageResizeSupportEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineStorageResizeSupport(input)
	return &out, nil
}

type StorageAutoGrowthSupport string

const (
	StorageAutoGrowthSupportDisabled StorageAutoGrowthSupport = "Disabled"
	StorageAutoGrowthSupportEnabled  StorageAutoGrowthSupport = "Enabled"
)

func PossibleValuesForStorageAutoGrowthSupport() []string {
	return []string{
		string(StorageAutoGrowthSupportDisabled),
		string(StorageAutoGrowthSupportEnabled),
	}
}

func (s *StorageAutoGrowthSupport) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageAutoGrowthSupport(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageAutoGrowthSupport(input string) (*StorageAutoGrowthSupport, error) {
	vals := map[string]StorageAutoGrowthSupport{
		"disabled": StorageAutoGrowthSupportDisabled,
		"enabled":  StorageAutoGrowthSupportEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageAutoGrowthSupport(input)
	return &out, nil
}

type ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupport string

const (
	ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupportDisabled ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupport = "Disabled"
	ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupportEnabled  ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupport = "Enabled"
)

func PossibleValuesForZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupport() []string {
	return []string{
		string(ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupportDisabled),
		string(ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupportEnabled),
	}
}

func (s *ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupport) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupport(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupport(input string) (*ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupport, error) {
	vals := map[string]ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupport{
		"disabled": ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupportDisabled,
		"enabled":  ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupportEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZoneRedundantHighAvailabilityAndGeographicallyRedundantBackupSupport(input)
	return &out, nil
}

type ZoneRedundantHighAvailabilitySupport string

const (
	ZoneRedundantHighAvailabilitySupportDisabled ZoneRedundantHighAvailabilitySupport = "Disabled"
	ZoneRedundantHighAvailabilitySupportEnabled  ZoneRedundantHighAvailabilitySupport = "Enabled"
)

func PossibleValuesForZoneRedundantHighAvailabilitySupport() []string {
	return []string{
		string(ZoneRedundantHighAvailabilitySupportDisabled),
		string(ZoneRedundantHighAvailabilitySupportEnabled),
	}
}

func (s *ZoneRedundantHighAvailabilitySupport) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZoneRedundantHighAvailabilitySupport(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZoneRedundantHighAvailabilitySupport(input string) (*ZoneRedundantHighAvailabilitySupport, error) {
	vals := map[string]ZoneRedundantHighAvailabilitySupport{
		"disabled": ZoneRedundantHighAvailabilitySupportDisabled,
		"enabled":  ZoneRedundantHighAvailabilitySupportEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZoneRedundantHighAvailabilitySupport(input)
	return &out, nil
}
