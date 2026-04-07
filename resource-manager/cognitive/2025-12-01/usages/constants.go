package usages

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaUsageStatus string

const (
	QuotaUsageStatusBlocked   QuotaUsageStatus = "Blocked"
	QuotaUsageStatusInOverage QuotaUsageStatus = "InOverage"
	QuotaUsageStatusIncluded  QuotaUsageStatus = "Included"
	QuotaUsageStatusUnknown   QuotaUsageStatus = "Unknown"
)

func PossibleValuesForQuotaUsageStatus() []string {
	return []string{
		string(QuotaUsageStatusBlocked),
		string(QuotaUsageStatusInOverage),
		string(QuotaUsageStatusIncluded),
		string(QuotaUsageStatusUnknown),
	}
}

func (s *QuotaUsageStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseQuotaUsageStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseQuotaUsageStatus(input string) (*QuotaUsageStatus, error) {
	vals := map[string]QuotaUsageStatus{
		"blocked":   QuotaUsageStatusBlocked,
		"inoverage": QuotaUsageStatusInOverage,
		"included":  QuotaUsageStatusIncluded,
		"unknown":   QuotaUsageStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := QuotaUsageStatus(input)
	return &out, nil
}

type UnitType string

const (
	UnitTypeBytes          UnitType = "Bytes"
	UnitTypeBytesPerSecond UnitType = "BytesPerSecond"
	UnitTypeCount          UnitType = "Count"
	UnitTypeCountPerSecond UnitType = "CountPerSecond"
	UnitTypeMilliseconds   UnitType = "Milliseconds"
	UnitTypePercent        UnitType = "Percent"
	UnitTypeSeconds        UnitType = "Seconds"
)

func PossibleValuesForUnitType() []string {
	return []string{
		string(UnitTypeBytes),
		string(UnitTypeBytesPerSecond),
		string(UnitTypeCount),
		string(UnitTypeCountPerSecond),
		string(UnitTypeMilliseconds),
		string(UnitTypePercent),
		string(UnitTypeSeconds),
	}
}

func (s *UnitType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnitType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnitType(input string) (*UnitType, error) {
	vals := map[string]UnitType{
		"bytes":          UnitTypeBytes,
		"bytespersecond": UnitTypeBytesPerSecond,
		"count":          UnitTypeCount,
		"countpersecond": UnitTypeCountPerSecond,
		"milliseconds":   UnitTypeMilliseconds,
		"percent":        UnitTypePercent,
		"seconds":        UnitTypeSeconds,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnitType(input)
	return &out, nil
}
