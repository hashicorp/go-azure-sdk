package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthReportRequest struct {
	AdditionalContext *string                     `json:"additionalContext,omitempty"`
	EvaluationRules   *HealthReportEvaluationRule `json:"evaluationRules,omitempty"`
	ExpiresInMinutes  *int64                      `json:"expiresInMinutes,omitempty"`
	HealthState       HealthState                 `json:"healthState"`
	SignalName        string                      `json:"signalName"`
	Value             *float64                    `json:"value,omitempty"`
}
