package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedApp struct {
	DeviceCount    *int64               `json:"deviceCount,omitempty"`
	DisplayName    *string              `json:"displayName,omitempty"`
	Id             *string              `json:"id,omitempty"`
	ManagedDevices *[]ManagedDevice     `json:"managedDevices,omitempty"`
	ODataType      *string              `json:"@odata.type,omitempty"`
	Platform       *DetectedAppPlatform `json:"platform,omitempty"`
	Publisher      *string              `json:"publisher,omitempty"`
	SizeInByte     *int64               `json:"sizeInByte,omitempty"`
	Version        *string              `json:"version,omitempty"`
}
