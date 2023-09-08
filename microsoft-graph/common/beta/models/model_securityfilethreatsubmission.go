package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileThreatSubmission struct {
	AdminReview     *SecuritySubmissionAdminReview            `json:"adminReview,omitempty"`
	Category        *SecurityFileThreatSubmissionCategory     `json:"category,omitempty"`
	ClientSource    *SecurityFileThreatSubmissionClientSource `json:"clientSource,omitempty"`
	ContentType     *SecurityFileThreatSubmissionContentType  `json:"contentType,omitempty"`
	CreatedBy       *SecuritySubmissionUserIdentity           `json:"createdBy,omitempty"`
	CreatedDateTime *string                                   `json:"createdDateTime,omitempty"`
	FileName        *string                                   `json:"fileName,omitempty"`
	Id              *string                                   `json:"id,omitempty"`
	ODataType       *string                                   `json:"@odata.type,omitempty"`
	Result          *SecuritySubmissionResult                 `json:"result,omitempty"`
	Source          *SecurityFileThreatSubmissionSource       `json:"source,omitempty"`
	Status          *SecurityFileThreatSubmissionStatus       `json:"status,omitempty"`
	TenantId        *string                                   `json:"tenantId,omitempty"`
}
