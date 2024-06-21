package costs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetCostProperties struct {
	CostThresholds     *[]CostThresholdProperties `json:"costThresholds,omitempty"`
	CycleEndDateTime   *string                    `json:"cycleEndDateTime,omitempty"`
	CycleStartDateTime *string                    `json:"cycleStartDateTime,omitempty"`
	CycleType          *ReportingCycleType        `json:"cycleType,omitempty"`
	Status             *TargetCostStatus          `json:"status,omitempty"`
	Target             *int64                     `json:"target,omitempty"`
}
