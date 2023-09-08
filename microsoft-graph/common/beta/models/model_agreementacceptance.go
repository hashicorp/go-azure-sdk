package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementAcceptance struct {
	AgreementFileId    *string                   `json:"agreementFileId,omitempty"`
	AgreementId        *string                   `json:"agreementId,omitempty"`
	DeviceDisplayName  *string                   `json:"deviceDisplayName,omitempty"`
	DeviceId           *string                   `json:"deviceId,omitempty"`
	DeviceOSType       *string                   `json:"deviceOSType,omitempty"`
	DeviceOSVersion    *string                   `json:"deviceOSVersion,omitempty"`
	ExpirationDateTime *string                   `json:"expirationDateTime,omitempty"`
	Id                 *string                   `json:"id,omitempty"`
	ODataType          *string                   `json:"@odata.type,omitempty"`
	RecordedDateTime   *string                   `json:"recordedDateTime,omitempty"`
	State              *AgreementAcceptanceState `json:"state,omitempty"`
	UserDisplayName    *string                   `json:"userDisplayName,omitempty"`
	UserEmail          *string                   `json:"userEmail,omitempty"`
	UserId             *string                   `json:"userId,omitempty"`
	UserPrincipalName  *string                   `json:"userPrincipalName,omitempty"`
}
