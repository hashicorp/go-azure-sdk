package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnrollmentProfile struct {
	AccountId                 *string                                                 `json:"accountId,omitempty"`
	ConfigureWifi             *bool                                                   `json:"configureWifi,omitempty"`
	CreatedDateTime           *string                                                 `json:"createdDateTime,omitempty"`
	Description               *string                                                 `json:"description,omitempty"`
	DisplayName               *string                                                 `json:"displayName,omitempty"`
	EnrolledDeviceCount       *int64                                                  `json:"enrolledDeviceCount,omitempty"`
	EnrollmentMode            *AndroidDeviceOwnerEnrollmentProfileEnrollmentMode      `json:"enrollmentMode,omitempty"`
	EnrollmentTokenType       *AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType `json:"enrollmentTokenType,omitempty"`
	EnrollmentTokenUsageCount *int64                                                  `json:"enrollmentTokenUsageCount,omitempty"`
	Id                        *string                                                 `json:"id,omitempty"`
	IsTeamsDeviceProfile      *bool                                                   `json:"isTeamsDeviceProfile,omitempty"`
	LastModifiedDateTime      *string                                                 `json:"lastModifiedDateTime,omitempty"`
	ODataType                 *string                                                 `json:"@odata.type,omitempty"`
	QrCodeContent             *string                                                 `json:"qrCodeContent,omitempty"`
	QrCodeImage               *MimeContent                                            `json:"qrCodeImage,omitempty"`
	RoleScopeTagIds           *[]string                                               `json:"roleScopeTagIds,omitempty"`
	TokenCreationDateTime     *string                                                 `json:"tokenCreationDateTime,omitempty"`
	TokenExpirationDateTime   *string                                                 `json:"tokenExpirationDateTime,omitempty"`
	TokenValue                *string                                                 `json:"tokenValue,omitempty"`
	WifiHidden                *bool                                                   `json:"wifiHidden,omitempty"`
	WifiPassword              *string                                                 `json:"wifiPassword,omitempty"`
	WifiSecurityType          *AndroidDeviceOwnerEnrollmentProfileWifiSecurityType    `json:"wifiSecurityType,omitempty"`
	WifiSsid                  *string                                                 `json:"wifiSsid,omitempty"`
}
