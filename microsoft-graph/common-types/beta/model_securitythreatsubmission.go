package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmission struct {
	AdminReview     *SecuritySubmissionAdminReview        `json:"adminReview,omitempty"`
	Category        *SecurityThreatSubmissionCategory     `json:"category,omitempty"`
	ClientSource    *SecurityThreatSubmissionClientSource `json:"clientSource,omitempty"`
	ContentType     *SecurityThreatSubmissionContentType  `json:"contentType,omitempty"`
	CreatedBy       *SecuritySubmissionUserIdentity       `json:"createdBy,omitempty"`
	CreatedDateTime *string                               `json:"createdDateTime,omitempty"`
	Id              *string                               `json:"id,omitempty"`
	ODataType       *string                               `json:"@odata.type,omitempty"`
	Result          *SecuritySubmissionResult             `json:"result,omitempty"`
	Source          *SecurityThreatSubmissionSource       `json:"source,omitempty"`
	Status          *SecurityThreatSubmissionStatus       `json:"status,omitempty"`
	TenantId        *string                               `json:"tenantId,omitempty"`
}
