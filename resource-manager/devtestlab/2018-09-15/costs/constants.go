package costs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostThresholdStatus string

const (
	CostThresholdStatusDisabled CostThresholdStatus = "Disabled"
	CostThresholdStatusEnabled  CostThresholdStatus = "Enabled"
)

func PossibleValuesForCostThresholdStatus() []string {
	return []string{
		string(CostThresholdStatusDisabled),
		string(CostThresholdStatusEnabled),
	}
}

func (s *CostThresholdStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCostThresholdStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCostThresholdStatus(input string) (*CostThresholdStatus, error) {
	vals := map[string]CostThresholdStatus{
		"disabled": CostThresholdStatusDisabled,
		"enabled":  CostThresholdStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CostThresholdStatus(input)
	return &out, nil
}

type CostType string

const (
	CostTypeProjected   CostType = "Projected"
	CostTypeReported    CostType = "Reported"
	CostTypeUnavailable CostType = "Unavailable"
)

func PossibleValuesForCostType() []string {
	return []string{
		string(CostTypeProjected),
		string(CostTypeReported),
		string(CostTypeUnavailable),
	}
}

func (s *CostType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCostType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCostType(input string) (*CostType, error) {
	vals := map[string]CostType{
		"projected":   CostTypeProjected,
		"reported":    CostTypeReported,
		"unavailable": CostTypeUnavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CostType(input)
	return &out, nil
}

type ReportingCycleType string

const (
	ReportingCycleTypeCalendarMonth ReportingCycleType = "CalendarMonth"
	ReportingCycleTypeCustom        ReportingCycleType = "Custom"
)

func PossibleValuesForReportingCycleType() []string {
	return []string{
		string(ReportingCycleTypeCalendarMonth),
		string(ReportingCycleTypeCustom),
	}
}

func (s *ReportingCycleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReportingCycleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReportingCycleType(input string) (*ReportingCycleType, error) {
	vals := map[string]ReportingCycleType{
		"calendarmonth": ReportingCycleTypeCalendarMonth,
		"custom":        ReportingCycleTypeCustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReportingCycleType(input)
	return &out, nil
}

type TargetCostStatus string

const (
	TargetCostStatusDisabled TargetCostStatus = "Disabled"
	TargetCostStatusEnabled  TargetCostStatus = "Enabled"
)

func PossibleValuesForTargetCostStatus() []string {
	return []string{
		string(TargetCostStatusDisabled),
		string(TargetCostStatusEnabled),
	}
}

func (s *TargetCostStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetCostStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetCostStatus(input string) (*TargetCostStatus, error) {
	vals := map[string]TargetCostStatus{
		"disabled": TargetCostStatusDisabled,
		"enabled":  TargetCostStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetCostStatus(input)
	return &out, nil
}
