package alertssuppressionrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertsSuppressionRuleProperties struct {
	AlertType              string                  `json:"alertType"`
	Comment                *string                 `json:"comment,omitempty"`
	ExpirationDateUtc      *string                 `json:"expirationDateUtc,omitempty"`
	LastModifiedUtc        *string                 `json:"lastModifiedUtc,omitempty"`
	Reason                 string                  `json:"reason"`
	State                  RuleState               `json:"state"`
	SuppressionAlertsScope *SuppressionAlertsScope `json:"suppressionAlertsScope,omitempty"`
}
