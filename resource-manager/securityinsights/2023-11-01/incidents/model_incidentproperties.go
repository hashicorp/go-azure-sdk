package incidents

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
	ProviderIncidentId     *string                       `json:"providerIncidentId,omitempty"`
	ProviderName           *string                       `json:"providerName,omitempty"`
	RelatedAnalyticRuleIds *[]string                     `json:"relatedAnalyticRuleIds,omitempty"`
	Severity               IncidentSeverity              `json:"severity"`
	Status                 IncidentStatus                `json:"status"`
	Title                  string                        `json:"title"`
}
