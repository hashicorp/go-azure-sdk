package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalLockConfiguration struct {
	AllProperties              *bool   `json:"allProperties,omitempty"`
	CredentialsWithUsageSign   *bool   `json:"credentialsWithUsageSign,omitempty"`
	CredentialsWithUsageVerify *bool   `json:"credentialsWithUsageVerify,omitempty"`
	IsEnabled                  *bool   `json:"isEnabled,omitempty"`
	ODataType                  *string `json:"@odata.type,omitempty"`
	TokenEncryptionKeyId       *bool   `json:"tokenEncryptionKeyId,omitempty"`
}
