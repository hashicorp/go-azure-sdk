package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplePushNotificationCertificate struct {
	AppleIdentifier                *string `json:"appleIdentifier,omitempty"`
	Certificate                    *string `json:"certificate,omitempty"`
	CertificateSerialNumber        *string `json:"certificateSerialNumber,omitempty"`
	CertificateUploadFailureReason *string `json:"certificateUploadFailureReason,omitempty"`
	CertificateUploadStatus        *string `json:"certificateUploadStatus,omitempty"`
	ExpirationDateTime             *string `json:"expirationDateTime,omitempty"`
	Id                             *string `json:"id,omitempty"`
	LastModifiedDateTime           *string `json:"lastModifiedDateTime,omitempty"`
	ODataType                      *string `json:"@odata.type,omitempty"`
	TopicIdentifier                *string `json:"topicIdentifier,omitempty"`
}
