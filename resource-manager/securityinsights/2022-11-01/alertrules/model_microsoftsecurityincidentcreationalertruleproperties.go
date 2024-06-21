package alertrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftSecurityIncidentCreationAlertRuleProperties struct {
	AlertRuleTemplateName     *string                      `json:"alertRuleTemplateName,omitempty"`
	Description               *string                      `json:"description,omitempty"`
	DisplayName               string                       `json:"displayName"`
	DisplayNamesExcludeFilter *[]string                    `json:"displayNamesExcludeFilter,omitempty"`
	DisplayNamesFilter        *[]string                    `json:"displayNamesFilter,omitempty"`
	Enabled                   bool                         `json:"enabled"`
	LastModifiedUtc           *string                      `json:"lastModifiedUtc,omitempty"`
	ProductFilter             MicrosoftSecurityProductName `json:"productFilter"`
	SeveritiesFilter          *[]AlertSeverity             `json:"severitiesFilter,omitempty"`
}
