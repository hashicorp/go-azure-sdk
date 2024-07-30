package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlThreatSubmission struct {
	AdminReview     *SecuritySubmissionAdminReview           `json:"adminReview,omitempty"`
	Category        *SecurityUrlThreatSubmissionCategory     `json:"category,omitempty"`
	ClientSource    *SecurityUrlThreatSubmissionClientSource `json:"clientSource,omitempty"`
	ContentType     *SecurityUrlThreatSubmissionContentType  `json:"contentType,omitempty"`
	CreatedBy       *SecuritySubmissionUserIdentity          `json:"createdBy,omitempty"`
	CreatedDateTime *string                                  `json:"createdDateTime,omitempty"`
	Id              *string                                  `json:"id,omitempty"`
	ODataType       *string                                  `json:"@odata.type,omitempty"`
	Result          *SecuritySubmissionResult                `json:"result,omitempty"`
	Source          *SecurityUrlThreatSubmissionSource       `json:"source,omitempty"`
	Status          *SecurityUrlThreatSubmissionStatus       `json:"status,omitempty"`
	TenantId        *string                                  `json:"tenantId,omitempty"`
	WebUrl          *string                                  `json:"webUrl,omitempty"`
}
