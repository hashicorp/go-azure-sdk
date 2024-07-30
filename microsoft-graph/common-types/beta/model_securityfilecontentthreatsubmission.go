package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileContentThreatSubmission struct {
	AdminReview     *SecuritySubmissionAdminReview                   `json:"adminReview,omitempty"`
	Category        *SecurityFileContentThreatSubmissionCategory     `json:"category,omitempty"`
	ClientSource    *SecurityFileContentThreatSubmissionClientSource `json:"clientSource,omitempty"`
	ContentType     *SecurityFileContentThreatSubmissionContentType  `json:"contentType,omitempty"`
	CreatedBy       *SecuritySubmissionUserIdentity                  `json:"createdBy,omitempty"`
	CreatedDateTime *string                                          `json:"createdDateTime,omitempty"`
	FileContent     *string                                          `json:"fileContent,omitempty"`
	FileName        *string                                          `json:"fileName,omitempty"`
	Id              *string                                          `json:"id,omitempty"`
	ODataType       *string                                          `json:"@odata.type,omitempty"`
	Result          *SecuritySubmissionResult                        `json:"result,omitempty"`
	Source          *SecurityFileContentThreatSubmissionSource       `json:"source,omitempty"`
	Status          *SecurityFileContentThreatSubmissionStatus       `json:"status,omitempty"`
	TenantId        *string                                          `json:"tenantId,omitempty"`
}
