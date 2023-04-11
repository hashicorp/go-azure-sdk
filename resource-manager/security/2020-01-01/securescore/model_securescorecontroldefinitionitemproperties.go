package securescore

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoreControlDefinitionItemProperties struct {
	AssessmentDefinitions *[]AzureResourceLink                `json:"assessmentDefinitions,omitempty"`
	Description           *string                             `json:"description,omitempty"`
	DisplayName           *string                             `json:"displayName,omitempty"`
	MaxScore              *int64                              `json:"maxScore,omitempty"`
	Source                *SecureScoreControlDefinitionSource `json:"source,omitempty"`
}
