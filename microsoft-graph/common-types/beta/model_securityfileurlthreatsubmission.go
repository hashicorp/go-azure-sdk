package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileUrlThreatSubmission struct {
	AdminReview     *SecuritySubmissionAdminReview               `json:"adminReview,omitempty"`
	Category        *SecurityFileUrlThreatSubmissionCategory     `json:"category,omitempty"`
	ClientSource    *SecurityFileUrlThreatSubmissionClientSource `json:"clientSource,omitempty"`
	ContentType     *SecurityFileUrlThreatSubmissionContentType  `json:"contentType,omitempty"`
	CreatedBy       *SecuritySubmissionUserIdentity              `json:"createdBy,omitempty"`
	CreatedDateTime *string                                      `json:"createdDateTime,omitempty"`
	FileName        *string                                      `json:"fileName,omitempty"`
	FileUrl         *string                                      `json:"fileUrl,omitempty"`
	Id              *string                                      `json:"id,omitempty"`
	ODataType       *string                                      `json:"@odata.type,omitempty"`
	Result          *SecuritySubmissionResult                    `json:"result,omitempty"`
	Source          *SecurityFileUrlThreatSubmissionSource       `json:"source,omitempty"`
	Status          *SecurityFileUrlThreatSubmissionStatus       `json:"status,omitempty"`
	TenantId        *string                                      `json:"tenantId,omitempty"`
}
