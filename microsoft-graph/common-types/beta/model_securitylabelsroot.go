package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityLabelsRoot struct {
	Authorities        *[]SecurityAuthorityTemplate         `json:"authorities,omitempty"`
	Categories         *[]SecurityCategoryTemplate          `json:"categories,omitempty"`
	Citations          *[]SecurityCitationTemplate          `json:"citations,omitempty"`
	Departments        *[]SecurityDepartmentTemplate        `json:"departments,omitempty"`
	FilePlanReferences *[]SecurityFilePlanReferenceTemplate `json:"filePlanReferences,omitempty"`
	Id                 *string                              `json:"id,omitempty"`
	ODataType          *string                              `json:"@odata.type,omitempty"`
	RetentionLabels    *[]SecurityRetentionLabel            `json:"retentionLabels,omitempty"`
}
