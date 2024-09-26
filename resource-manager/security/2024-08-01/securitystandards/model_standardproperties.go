package securitystandards

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardProperties struct {
	Assessments           *[]PartialAssessmentProperties `json:"assessments,omitempty"`
	CloudProviders        *[]StandardSupportedCloud      `json:"cloudProviders,omitempty"`
	Description           *string                        `json:"description,omitempty"`
	DisplayName           *string                        `json:"displayName,omitempty"`
	Metadata              *StandardMetadata              `json:"metadata,omitempty"`
	PolicySetDefinitionId *string                        `json:"policySetDefinitionId,omitempty"`
	StandardType          *StandardType                  `json:"standardType,omitempty"`
}
