package alertrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FusionAlertRuleProperties struct {
	AlertRuleTemplateName     string                            `json:"alertRuleTemplateName"`
	Description               *string                           `json:"description,omitempty"`
	DisplayName               *string                           `json:"displayName,omitempty"`
	Enabled                   bool                              `json:"enabled"`
	LastModifiedUtc           *string                           `json:"lastModifiedUtc,omitempty"`
	ScenarioExclusionPatterns *[]FusionScenarioExclusionPattern `json:"scenarioExclusionPatterns,omitempty"`
	Severity                  *AlertSeverity                    `json:"severity,omitempty"`
	SourceSettings            *[]FusionSourceSettings           `json:"sourceSettings,omitempty"`
	Tactics                   *[]AttackTactic                   `json:"tactics,omitempty"`
	Techniques                *[]string                         `json:"techniques,omitempty"`
}
