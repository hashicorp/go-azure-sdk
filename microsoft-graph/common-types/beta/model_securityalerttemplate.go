package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertTemplate struct {
	Category           *string                        `json:"category,omitempty"`
	Description        *string                        `json:"description,omitempty"`
	ImpactedAssets     *[]SecurityImpactedAsset       `json:"impactedAssets,omitempty"`
	MitreTechniques    *[]string                      `json:"mitreTechniques,omitempty"`
	ODataType          *string                        `json:"@odata.type,omitempty"`
	RecommendedActions *string                        `json:"recommendedActions,omitempty"`
	Severity           *SecurityAlertTemplateSeverity `json:"severity,omitempty"`
	Title              *string                        `json:"title,omitempty"`
}
