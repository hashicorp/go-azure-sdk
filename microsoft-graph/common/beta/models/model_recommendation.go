package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Recommendation struct {
	ActionSteps           *[]ActionStep                     `json:"actionSteps,omitempty"`
	Benefits              *string                           `json:"benefits,omitempty"`
	Category              *RecommendationCategory           `json:"category,omitempty"`
	CreatedDateTime       *string                           `json:"createdDateTime,omitempty"`
	DisplayName           *string                           `json:"displayName,omitempty"`
	FeatureAreas          *[]RecommendationFeatureAreas     `json:"featureAreas,omitempty"`
	Id                    *string                           `json:"id,omitempty"`
	ImpactStartDateTime   *string                           `json:"impactStartDateTime,omitempty"`
	ImpactType            *string                           `json:"impactType,omitempty"`
	ImpactedResources     *[]ImpactedResource               `json:"impactedResources,omitempty"`
	Insights              *string                           `json:"insights,omitempty"`
	LastCheckedDateTime   *string                           `json:"lastCheckedDateTime,omitempty"`
	LastModifiedBy        *string                           `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime  *string                           `json:"lastModifiedDateTime,omitempty"`
	ODataType             *string                           `json:"@odata.type,omitempty"`
	PostponeUntilDateTime *string                           `json:"postponeUntilDateTime,omitempty"`
	Priority              *RecommendationPriority           `json:"priority,omitempty"`
	RecommendationType    *RecommendationRecommendationType `json:"recommendationType,omitempty"`
	RemediationImpact     *string                           `json:"remediationImpact,omitempty"`
	Status                *RecommendationStatus             `json:"status,omitempty"`
}
