package alertruletemplates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduledAlertRuleTemplateProperties struct {
	AlertDetailsOverride             *AlertDetailsOverride          `json:"alertDetailsOverride,omitempty"`
	AlertRulesCreatedByTemplateCount *int64                         `json:"alertRulesCreatedByTemplateCount,omitempty"`
	CreatedDateUTC                   *string                        `json:"createdDateUTC,omitempty"`
	CustomDetails                    *map[string]string             `json:"customDetails,omitempty"`
	Description                      *string                        `json:"description,omitempty"`
	DisplayName                      *string                        `json:"displayName,omitempty"`
	EntityMappings                   *[]EntityMapping               `json:"entityMappings,omitempty"`
	EventGroupingSettings            *EventGroupingSettings         `json:"eventGroupingSettings,omitempty"`
	LastUpdatedDateUTC               *string                        `json:"lastUpdatedDateUTC,omitempty"`
	Query                            *string                        `json:"query,omitempty"`
	QueryFrequency                   *string                        `json:"queryFrequency,omitempty"`
	QueryPeriod                      *string                        `json:"queryPeriod,omitempty"`
	RequiredDataConnectors           *[]AlertRuleTemplateDataSource `json:"requiredDataConnectors,omitempty"`
	Severity                         *AlertSeverity                 `json:"severity,omitempty"`
	Status                           *TemplateStatus                `json:"status,omitempty"`
	Tactics                          *[]AttackTactic                `json:"tactics,omitempty"`
	Techniques                       *[]string                      `json:"techniques,omitempty"`
	TriggerOperator                  *TriggerOperator               `json:"triggerOperator,omitempty"`
	TriggerThreshold                 *int64                         `json:"triggerThreshold,omitempty"`
	Version                          *string                        `json:"version,omitempty"`
}
