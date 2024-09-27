package customrecommendations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomRecommendationProperties struct {
	AssessmentKey          *string                          `json:"assessmentKey,omitempty"`
	CloudProviders         *[]RecommendationSupportedClouds `json:"cloudProviders,omitempty"`
	Description            *string                          `json:"description,omitempty"`
	DisplayName            *string                          `json:"displayName,omitempty"`
	Query                  *string                          `json:"query,omitempty"`
	RemediationDescription *string                          `json:"remediationDescription,omitempty"`
	SecurityIssue          *SecurityIssue                   `json:"securityIssue,omitempty"`
	Severity               *SeverityEnum                    `json:"severity,omitempty"`
}
