package replicationvaulthealth

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthErrorSummary struct {
	AffectedResourceCorrelationIds *[]string            `json:"affectedResourceCorrelationIds,omitempty"`
	AffectedResourceSubtype        *string              `json:"affectedResourceSubtype,omitempty"`
	AffectedResourceType           *string              `json:"affectedResourceType,omitempty"`
	Category                       *HealthErrorCategory `json:"category,omitempty"`
	Severity                       *Severity            `json:"severity,omitempty"`
	SummaryCode                    *string              `json:"summaryCode,omitempty"`
	SummaryMessage                 *string              `json:"summaryMessage,omitempty"`
}
