package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedMessageEvidence struct {
	AntiSpamDirection        *string                                           `json:"antiSpamDirection,omitempty"`
	AttachmentsCount         *int64                                            `json:"attachmentsCount,omitempty"`
	CreatedDateTime          *string                                           `json:"createdDateTime,omitempty"`
	DeliveryAction           *string                                           `json:"deliveryAction,omitempty"`
	DeliveryLocation         *string                                           `json:"deliveryLocation,omitempty"`
	DetailedRoles            *[]string                                         `json:"detailedRoles,omitempty"`
	InternetMessageId        *string                                           `json:"internetMessageId,omitempty"`
	Language                 *string                                           `json:"language,omitempty"`
	NetworkMessageId         *string                                           `json:"networkMessageId,omitempty"`
	ODataType                *string                                           `json:"@odata.type,omitempty"`
	P1Sender                 *SecurityEmailSender                              `json:"p1Sender,omitempty"`
	P2Sender                 *SecurityEmailSender                              `json:"p2Sender,omitempty"`
	ReceivedDateTime         *string                                           `json:"receivedDateTime,omitempty"`
	RecipientEmailAddress    *string                                           `json:"recipientEmailAddress,omitempty"`
	RemediationStatus        *SecurityAnalyzedMessageEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                           `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityAnalyzedMessageEvidenceRoles             `json:"roles,omitempty"`
	SenderIp                 *string                                           `json:"senderIp,omitempty"`
	Subject                  *string                                           `json:"subject,omitempty"`
	Tags                     *[]string                                         `json:"tags,omitempty"`
	ThreatDetectionMethods   *[]string                                         `json:"threatDetectionMethods,omitempty"`
	Threats                  *[]string                                         `json:"threats,omitempty"`
	UrlCount                 *int64                                            `json:"urlCount,omitempty"`
	Urls                     *[]string                                         `json:"urls,omitempty"`
	Urn                      *string                                           `json:"urn,omitempty"`
	Verdict                  *SecurityAnalyzedMessageEvidenceVerdict           `json:"verdict,omitempty"`
}
