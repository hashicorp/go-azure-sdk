package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnrollmentProfile struct {
	AccountId               *string      `json:"accountId,omitempty"`
	CreatedDateTime         *string      `json:"createdDateTime,omitempty"`
	Description             *string      `json:"description,omitempty"`
	DisplayName             *string      `json:"displayName,omitempty"`
	EnrolledDeviceCount     *int64       `json:"enrolledDeviceCount,omitempty"`
	Id                      *string      `json:"id,omitempty"`
	LastModifiedDateTime    *string      `json:"lastModifiedDateTime,omitempty"`
	ODataType               *string      `json:"@odata.type,omitempty"`
	QrCodeContent           *string      `json:"qrCodeContent,omitempty"`
	QrCodeImage             *MimeContent `json:"qrCodeImage,omitempty"`
	TokenExpirationDateTime *string      `json:"tokenExpirationDateTime,omitempty"`
	TokenValue              *string      `json:"tokenValue,omitempty"`
}
