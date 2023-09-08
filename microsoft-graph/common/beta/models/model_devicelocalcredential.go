package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceLocalCredential struct {
	AccountName    *string `json:"accountName,omitempty"`
	AccountSid     *string `json:"accountSid,omitempty"`
	BackupDateTime *string `json:"backupDateTime,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
	PasswordBase64 *string `json:"passwordBase64,omitempty"`
}
