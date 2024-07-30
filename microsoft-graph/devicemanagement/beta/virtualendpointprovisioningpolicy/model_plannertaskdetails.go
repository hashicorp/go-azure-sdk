package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskDetails struct {
	Checklist              *PlannerChecklistItems                   `json:"checklist,omitempty"`
	CompletionRequirements *PlannerTaskCompletionRequirementDetails `json:"completionRequirements,omitempty"`
	Description            *string                                  `json:"description,omitempty"`
	Id                     *string                                  `json:"id,omitempty"`
	Notes                  *ItemBody                                `json:"notes,omitempty"`
	ODataType              *string                                  `json:"@odata.type,omitempty"`
	PreviewType            *PlannerTaskDetailsPreviewType           `json:"previewType,omitempty"`
	References             *PlannerExternalReferences               `json:"references,omitempty"`
}
