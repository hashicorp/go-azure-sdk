package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateState struct {
	DeviceDisplayName    *string                   `json:"deviceDisplayName,omitempty"`
	DeviceId             *string                   `json:"deviceId,omitempty"`
	FeatureUpdateVersion *string                   `json:"featureUpdateVersion,omitempty"`
	Id                   *string                   `json:"id,omitempty"`
	LastScanDateTime     *string                   `json:"lastScanDateTime,omitempty"`
	LastSyncDateTime     *string                   `json:"lastSyncDateTime,omitempty"`
	ODataType            *string                   `json:"@odata.type,omitempty"`
	QualityUpdateVersion *string                   `json:"qualityUpdateVersion,omitempty"`
	Status               *WindowsUpdateStateStatus `json:"status,omitempty"`
	UserId               *string                   `json:"userId,omitempty"`
	UserPrincipalName    *string                   `json:"userPrincipalName,omitempty"`
}
