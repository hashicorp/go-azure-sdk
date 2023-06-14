package locationcapabilities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityGroup string

const (
	CapabilityGroupSupportedEditions                CapabilityGroup = "supportedEditions"
	CapabilityGroupSupportedElasticPoolEditions     CapabilityGroup = "supportedElasticPoolEditions"
	CapabilityGroupSupportedInstancePoolEditions    CapabilityGroup = "supportedInstancePoolEditions"
	CapabilityGroupSupportedManagedInstanceEditions CapabilityGroup = "supportedManagedInstanceEditions"
	CapabilityGroupSupportedManagedInstanceVersions CapabilityGroup = "supportedManagedInstanceVersions"
)

func PossibleValuesForCapabilityGroup() []string {
	return []string{
		string(CapabilityGroupSupportedEditions),
		string(CapabilityGroupSupportedElasticPoolEditions),
		string(CapabilityGroupSupportedInstancePoolEditions),
		string(CapabilityGroupSupportedManagedInstanceEditions),
		string(CapabilityGroupSupportedManagedInstanceVersions),
	}
}

func (s *CapabilityGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCapabilityGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCapabilityGroup(input string) (*CapabilityGroup, error) {
	vals := map[string]CapabilityGroup{
		"supportededitions":                CapabilityGroupSupportedEditions,
		"supportedelasticpooleditions":     CapabilityGroupSupportedElasticPoolEditions,
		"supportedinstancepooleditions":    CapabilityGroupSupportedInstancePoolEditions,
		"supportedmanagedinstanceeditions": CapabilityGroupSupportedManagedInstanceEditions,
		"supportedmanagedinstanceversions": CapabilityGroupSupportedManagedInstanceVersions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CapabilityGroup(input)
	return &out, nil
}

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

type LogSizeUnit string

const (
	LogSizeUnitGigabytes LogSizeUnit = "Gigabytes"
	LogSizeUnitMegabytes LogSizeUnit = "Megabytes"
	LogSizeUnitPercent   LogSizeUnit = "Percent"
	LogSizeUnitPetabytes LogSizeUnit = "Petabytes"
	LogSizeUnitTerabytes LogSizeUnit = "Terabytes"
)

func PossibleValuesForLogSizeUnit() []string {
	return []string{
		string(LogSizeUnitGigabytes),
		string(LogSizeUnitMegabytes),
		string(LogSizeUnitPercent),
		string(LogSizeUnitPetabytes),
		string(LogSizeUnitTerabytes),
	}
}

func (s *LogSizeUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLogSizeUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLogSizeUnit(input string) (*LogSizeUnit, error) {
	vals := map[string]LogSizeUnit{
		"gigabytes": LogSizeUnitGigabytes,
		"megabytes": LogSizeUnitMegabytes,
		"percent":   LogSizeUnitPercent,
		"petabytes": LogSizeUnitPetabytes,
		"terabytes": LogSizeUnitTerabytes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogSizeUnit(input)
	return &out, nil
}

type MaxSizeUnit string

const (
	MaxSizeUnitGigabytes MaxSizeUnit = "Gigabytes"
	MaxSizeUnitMegabytes MaxSizeUnit = "Megabytes"
	MaxSizeUnitPetabytes MaxSizeUnit = "Petabytes"
	MaxSizeUnitTerabytes MaxSizeUnit = "Terabytes"
)

func PossibleValuesForMaxSizeUnit() []string {
	return []string{
		string(MaxSizeUnitGigabytes),
		string(MaxSizeUnitMegabytes),
		string(MaxSizeUnitPetabytes),
		string(MaxSizeUnitTerabytes),
	}
}

func (s *MaxSizeUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMaxSizeUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMaxSizeUnit(input string) (*MaxSizeUnit, error) {
	vals := map[string]MaxSizeUnit{
		"gigabytes": MaxSizeUnitGigabytes,
		"megabytes": MaxSizeUnitMegabytes,
		"petabytes": MaxSizeUnitPetabytes,
		"terabytes": MaxSizeUnitTerabytes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MaxSizeUnit(input)
	return &out, nil
}

type PauseDelayTimeUnit string

const (
	PauseDelayTimeUnitMinutes PauseDelayTimeUnit = "Minutes"
)

func PossibleValuesForPauseDelayTimeUnit() []string {
	return []string{
		string(PauseDelayTimeUnitMinutes),
	}
}

func (s *PauseDelayTimeUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePauseDelayTimeUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePauseDelayTimeUnit(input string) (*PauseDelayTimeUnit, error) {
	vals := map[string]PauseDelayTimeUnit{
		"minutes": PauseDelayTimeUnitMinutes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PauseDelayTimeUnit(input)
	return &out, nil
}

type PerformanceLevelUnit string

const (
	PerformanceLevelUnitDTU    PerformanceLevelUnit = "DTU"
	PerformanceLevelUnitVCores PerformanceLevelUnit = "VCores"
)

func PossibleValuesForPerformanceLevelUnit() []string {
	return []string{
		string(PerformanceLevelUnitDTU),
		string(PerformanceLevelUnitVCores),
	}
}

func (s *PerformanceLevelUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePerformanceLevelUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePerformanceLevelUnit(input string) (*PerformanceLevelUnit, error) {
	vals := map[string]PerformanceLevelUnit{
		"dtu":    PerformanceLevelUnitDTU,
		"vcores": PerformanceLevelUnitVCores,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PerformanceLevelUnit(input)
	return &out, nil
}

type StorageAccountType string

const (
	StorageAccountTypeGRS StorageAccountType = "GRS"
	StorageAccountTypeLRS StorageAccountType = "LRS"
	StorageAccountTypeZRS StorageAccountType = "ZRS"
)

func PossibleValuesForStorageAccountType() []string {
	return []string{
		string(StorageAccountTypeGRS),
		string(StorageAccountTypeLRS),
		string(StorageAccountTypeZRS),
	}
}

func (s *StorageAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageAccountType(input string) (*StorageAccountType, error) {
	vals := map[string]StorageAccountType{
		"grs": StorageAccountTypeGRS,
		"lrs": StorageAccountTypeLRS,
		"zrs": StorageAccountTypeZRS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageAccountType(input)
	return &out, nil
}
