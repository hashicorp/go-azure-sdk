package automationrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationRuleModifyPropertiesActionActionConfiguration struct {
	Classification        *IncidentClassification       `json:"classification,omitempty"`
	ClassificationComment *string                       `json:"classificationComment,omitempty"`
	ClassificationReason  *IncidentClassificationReason `json:"classificationReason,omitempty"`
	Labels                *[]IncidentLabel              `json:"labels,omitempty"`
	Owner                 *IncidentOwnerInfo            `json:"owner"`
	Severity              *IncidentSeverity             `json:"severity,omitempty"`
	Status                *IncidentStatus               `json:"status,omitempty"`
}
