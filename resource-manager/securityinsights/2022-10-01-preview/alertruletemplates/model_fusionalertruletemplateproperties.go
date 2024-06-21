package alertruletemplates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FusionAlertRuleTemplateProperties struct {
	AlertRulesCreatedByTemplateCount *int64                         `json:"alertRulesCreatedByTemplateCount,omitempty"`
	CreatedDateUTC                   *string                        `json:"createdDateUTC,omitempty"`
	Description                      *string                        `json:"description,omitempty"`
	DisplayName                      *string                        `json:"displayName,omitempty"`
	LastUpdatedDateUTC               *string                        `json:"lastUpdatedDateUTC,omitempty"`
	RequiredDataConnectors           *[]AlertRuleTemplateDataSource `json:"requiredDataConnectors,omitempty"`
	Severity                         *AlertSeverity                 `json:"severity,omitempty"`
	SourceSettings                   *[]FusionTemplateSourceSetting `json:"sourceSettings,omitempty"`
	Status                           *TemplateStatus                `json:"status,omitempty"`
	Tactics                          *[]AttackTactic                `json:"tactics,omitempty"`
	Techniques                       *[]string                      `json:"techniques,omitempty"`
}
