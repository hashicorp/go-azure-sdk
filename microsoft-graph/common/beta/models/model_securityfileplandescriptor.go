package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFilePlanDescriptor struct {
	Authority                 *SecurityAuthority                 `json:"authority,omitempty"`
	AuthorityTemplate         *SecurityAuthorityTemplate         `json:"authorityTemplate,omitempty"`
	Category                  *SecurityAppliedCategory           `json:"category,omitempty"`
	CategoryTemplate          *SecurityCategoryTemplate          `json:"categoryTemplate,omitempty"`
	Citation                  *SecurityCitation                  `json:"citation,omitempty"`
	CitationTemplate          *SecurityCitationTemplate          `json:"citationTemplate,omitempty"`
	Department                *SecurityDepartment                `json:"department,omitempty"`
	DepartmentTemplate        *SecurityDepartmentTemplate        `json:"departmentTemplate,omitempty"`
	FilePlanReference         *SecurityFilePlanReference         `json:"filePlanReference,omitempty"`
	FilePlanReferenceTemplate *SecurityFilePlanReferenceTemplate `json:"filePlanReferenceTemplate,omitempty"`
	Id                        *string                            `json:"id,omitempty"`
	ODataType                 *string                            `json:"@odata.type,omitempty"`
}
