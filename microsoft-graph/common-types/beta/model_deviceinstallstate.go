package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceInstallState struct {
	DeviceId         *string                         `json:"deviceId,omitempty"`
	DeviceName       *string                         `json:"deviceName,omitempty"`
	ErrorCode        *string                         `json:"errorCode,omitempty"`
	Id               *string                         `json:"id,omitempty"`
	InstallState     *DeviceInstallStateInstallState `json:"installState,omitempty"`
	LastSyncDateTime *string                         `json:"lastSyncDateTime,omitempty"`
	ODataType        *string                         `json:"@odata.type,omitempty"`
	OsDescription    *string                         `json:"osDescription,omitempty"`
	OsVersion        *string                         `json:"osVersion,omitempty"`
	UserName         *string                         `json:"userName,omitempty"`
}
