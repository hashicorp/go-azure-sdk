package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceLocalCredentialInfo struct {
	Credentials        *[]DeviceLocalCredential `json:"credentials,omitempty"`
	DeviceName         *string                  `json:"deviceName,omitempty"`
	Id                 *string                  `json:"id,omitempty"`
	LastBackupDateTime *string                  `json:"lastBackupDateTime,omitempty"`
	ODataType          *string                  `json:"@odata.type,omitempty"`
	RefreshDateTime    *string                  `json:"refreshDateTime,omitempty"`
}
