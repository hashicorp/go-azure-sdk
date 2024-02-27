package incidents

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentProperties struct {
	AdditionalData         *IncidentAdditionalData       `json:"additionalData,omitempty"`
	Classification         *IncidentClassification       `json:"classification,omitempty"`
	ClassificationComment  *string                       `json:"classificationComment,omitempty"`
	ClassificationReason   *IncidentClassificationReason `json:"classificationReason,omitempty"`
	CreatedTimeUtc         *string                       `json:"createdTimeUtc,omitempty"`
	Description            *string                       `json:"description,omitempty"`
	FirstActivityTimeUtc   *string                       `json:"firstActivityTimeUtc,omitempty"`
	IncidentNumber         *int64                        `json:"incidentNumber,omitempty"`
	IncidentUrl            *string                       `json:"incidentUrl,omitempty"`
	Labels                 *[]IncidentLabel              `json:"labels,omitempty"`
	LastActivityTimeUtc    *string                       `json:"lastActivityTimeUtc,omitempty"`
	LastModifiedTimeUtc    *string                       `json:"lastModifiedTimeUtc,omitempty"`
	Owner                  *IncidentOwnerInfo            `json:"owner,omitempty"`
	RelatedAnalyticRuleIds *[]string                     `json:"relatedAnalyticRuleIds,omitempty"`
	Severity               IncidentSeverity              `json:"severity"`
	Status                 IncidentStatus                `json:"status"`
	Title                  string                        `json:"title"`
}

func (o *IncidentProperties) GetCreatedTimeUtcAsTime() (*time.Time, error) {
	if o.CreatedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *IncidentProperties) GetFirstActivityTimeUtcAsTime() (*time.Time, error) {
	if o.FirstActivityTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.FirstActivityTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *IncidentProperties) SetFirstActivityTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.FirstActivityTimeUtc = &formatted
}

func (o *IncidentProperties) GetLastActivityTimeUtcAsTime() (*time.Time, error) {
	if o.LastActivityTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastActivityTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *IncidentProperties) SetLastActivityTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastActivityTimeUtc = &formatted
}

func (o *IncidentProperties) GetLastModifiedTimeUtcAsTime() (*time.Time, error) {
	if o.LastModifiedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedTimeUtc, "2006-01-02T15:04:05Z07:00")
}
