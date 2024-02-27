package alertrules

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryBasedAlertRuleProperties struct {
	AlertDetailsOverride  *AlertDetailsOverride  `json:"alertDetailsOverride,omitempty"`
	AlertRuleTemplateName *string                `json:"alertRuleTemplateName,omitempty"`
	CustomDetails         *map[string]string     `json:"customDetails,omitempty"`
	Description           *string                `json:"description,omitempty"`
	DisplayName           string                 `json:"displayName"`
	Enabled               bool                   `json:"enabled"`
	EntityMappings        *[]EntityMapping       `json:"entityMappings,omitempty"`
	IncidentConfiguration *IncidentConfiguration `json:"incidentConfiguration,omitempty"`
	LastModifiedUtc       *string                `json:"lastModifiedUtc,omitempty"`
	Query                 *string                `json:"query,omitempty"`
	Severity              *AlertSeverity         `json:"severity,omitempty"`
	SuppressionDuration   string                 `json:"suppressionDuration"`
	SuppressionEnabled    bool                   `json:"suppressionEnabled"`
	Tactics               *[]AttackTactic        `json:"tactics,omitempty"`
	TemplateVersion       *string                `json:"templateVersion,omitempty"`
}

func (o *QueryBasedAlertRuleProperties) GetLastModifiedUtcAsTime() (*time.Time, error) {
	if o.LastModifiedUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *QueryBasedAlertRuleProperties) SetLastModifiedUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedUtc = &formatted
}
