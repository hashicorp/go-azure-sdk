package costs

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
