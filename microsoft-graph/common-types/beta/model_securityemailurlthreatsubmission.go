package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmission struct {
	AdminReview                  *SecuritySubmissionAdminReview                    `json:"adminReview,omitempty"`
	AttackSimulationInfo         *SecurityAttackSimulationInfo                     `json:"attackSimulationInfo,omitempty"`
	Category                     *SecurityEmailUrlThreatSubmissionCategory         `json:"category,omitempty"`
	ClientSource                 *SecurityEmailUrlThreatSubmissionClientSource     `json:"clientSource,omitempty"`
	ContentType                  *SecurityEmailUrlThreatSubmissionContentType      `json:"contentType,omitempty"`
	CreatedBy                    *SecuritySubmissionUserIdentity                   `json:"createdBy,omitempty"`
	CreatedDateTime              *string                                           `json:"createdDateTime,omitempty"`
	Id                           *string                                           `json:"id,omitempty"`
	InternetMessageId            *string                                           `json:"internetMessageId,omitempty"`
	MessageUrl                   *string                                           `json:"messageUrl,omitempty"`
	ODataType                    *string                                           `json:"@odata.type,omitempty"`
	OriginalCategory             *SecurityEmailUrlThreatSubmissionOriginalCategory `json:"originalCategory,omitempty"`
	ReceivedDateTime             *string                                           `json:"receivedDateTime,omitempty"`
	RecipientEmailAddress        *string                                           `json:"recipientEmailAddress,omitempty"`
	Result                       *SecuritySubmissionResult                         `json:"result,omitempty"`
	Sender                       *string                                           `json:"sender,omitempty"`
	SenderIP                     *string                                           `json:"senderIP,omitempty"`
	Source                       *SecurityEmailUrlThreatSubmissionSource           `json:"source,omitempty"`
	Status                       *SecurityEmailUrlThreatSubmissionStatus           `json:"status,omitempty"`
	Subject                      *string                                           `json:"subject,omitempty"`
	TenantAllowOrBlockListAction *SecurityTenantAllowOrBlockListAction             `json:"tenantAllowOrBlockListAction,omitempty"`
	TenantId                     *string                                           `json:"tenantId,omitempty"`
}
