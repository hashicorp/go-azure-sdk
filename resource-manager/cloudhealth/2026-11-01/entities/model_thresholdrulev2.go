package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThresholdRuleV2 struct {
	Operator    SignalOperator               `json:"operator"`
	Sensitivity *DynamicThresholdSensitivity `json:"sensitivity,omitempty"`
	Threshold   *float64                     `json:"threshold,omitempty"`
}
