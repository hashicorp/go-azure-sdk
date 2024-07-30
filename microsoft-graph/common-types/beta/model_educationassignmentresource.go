package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentResource struct {
	DependentResources       *[]EducationAssignmentResource `json:"dependentResources,omitempty"`
	DistributeForStudentWork *bool                          `json:"distributeForStudentWork,omitempty"`
	Id                       *string                        `json:"id,omitempty"`
	ODataType                *string                        `json:"@odata.type,omitempty"`
	Resource                 *EducationResource             `json:"resource,omitempty"`
}
