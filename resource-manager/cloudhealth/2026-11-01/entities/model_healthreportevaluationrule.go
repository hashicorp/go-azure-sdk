package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthReportEvaluationRule struct {
	DegradedRule  *ThresholdRuleV2 `json:"degradedRule,omitempty"`
	UnhealthyRule ThresholdRuleV2  `json:"unhealthyRule"`
}
