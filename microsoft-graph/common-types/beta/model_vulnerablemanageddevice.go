package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VulnerableManagedDevice struct {
	DisplayName      *string `json:"displayName,omitempty"`
	Id               *string `json:"id,omitempty"`
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`
	ManagedDeviceId  *string `json:"managedDeviceId,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
}
