package alertssuppressionrules

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *AlertsSuppressionRuleProperties) GetExpirationDateUtcAsTime() (*time.Time, error) {
	if o.ExpirationDateUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpirationDateUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *AlertsSuppressionRuleProperties) SetExpirationDateUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpirationDateUtc = &formatted
}

func (o *AlertsSuppressionRuleProperties) GetLastModifiedUtcAsTime() (*time.Time, error) {
	if o.LastModifiedUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedUtc, "2006-01-02T15:04:05Z07:00")
}
