package costs

import "strings"

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
