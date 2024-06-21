package getrecommendations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationProperties struct {
	Actions                   *[]map[string]interface{} `json:"actions,omitempty"`
	Category                  *Category                 `json:"category,omitempty"`
	Description               *string                   `json:"description,omitempty"`
	ExposedMetadataProperties *map[string]interface{}   `json:"exposedMetadataProperties,omitempty"`
	ExtendedProperties        *map[string]string        `json:"extendedProperties,omitempty"`
	Impact                    *Impact                   `json:"impact,omitempty"`
	ImpactedField             *string                   `json:"impactedField,omitempty"`
	ImpactedValue             *string                   `json:"impactedValue,omitempty"`
	Label                     *string                   `json:"label,omitempty"`
	LastUpdated               *string                   `json:"lastUpdated,omitempty"`
	LearnMoreLink             *string                   `json:"learnMoreLink,omitempty"`
	Metadata                  *map[string]interface{}   `json:"metadata,omitempty"`
	PotentialBenefits         *string                   `json:"potentialBenefits,omitempty"`
	RecommendationTypeId      *string                   `json:"recommendationTypeId,omitempty"`
	Remediation               *map[string]interface{}   `json:"remediation,omitempty"`
	ResourceMetadata          *ResourceMetadata         `json:"resourceMetadata,omitempty"`
	ShortDescription          *ShortDescription         `json:"shortDescription,omitempty"`
	SuppressionIds            *[]string                 `json:"suppressionIds,omitempty"`
}
