package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoPauseDelayTimeRange struct {
	Default         *int64              `json:"default,omitempty"`
	DoNotPauseValue *int64              `json:"doNotPauseValue,omitempty"`
	MaxValue        *int64              `json:"maxValue,omitempty"`
	MinValue        *int64              `json:"minValue,omitempty"`
	StepSize        *int64              `json:"stepSize,omitempty"`
	Unit            *PauseDelayTimeUnit `json:"unit,omitempty"`
}
