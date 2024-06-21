package alertruletemplates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftSecurityIncidentCreationAlertRuleTemplateProperties struct {
	AlertRulesCreatedByTemplateCount *int64                         `json:"alertRulesCreatedByTemplateCount,omitempty"`
	CreatedDateUTC                   *string                        `json:"createdDateUTC,omitempty"`
	Description                      *string                        `json:"description,omitempty"`
	DisplayName                      *string                        `json:"displayName,omitempty"`
	DisplayNamesExcludeFilter        *[]string                      `json:"displayNamesExcludeFilter,omitempty"`
	DisplayNamesFilter               *[]string                      `json:"displayNamesFilter,omitempty"`
	LastUpdatedDateUTC               *string                        `json:"lastUpdatedDateUTC,omitempty"`
	ProductFilter                    *MicrosoftSecurityProductName  `json:"productFilter,omitempty"`
	RequiredDataConnectors           *[]AlertRuleTemplateDataSource `json:"requiredDataConnectors,omitempty"`
	SeveritiesFilter                 *[]AlertSeverity               `json:"severitiesFilter,omitempty"`
	Status                           *TemplateStatus                `json:"status,omitempty"`
}
