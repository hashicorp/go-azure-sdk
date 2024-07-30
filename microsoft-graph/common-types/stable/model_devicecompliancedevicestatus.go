package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceDeviceStatus struct {
	ComplianceGracePeriodExpirationDateTime *string                             `json:"complianceGracePeriodExpirationDateTime,omitempty"`
	DeviceDisplayName                       *string                             `json:"deviceDisplayName,omitempty"`
	DeviceModel                             *string                             `json:"deviceModel,omitempty"`
	Id                                      *string                             `json:"id,omitempty"`
	LastReportedDateTime                    *string                             `json:"lastReportedDateTime,omitempty"`
	ODataType                               *string                             `json:"@odata.type,omitempty"`
	Status                                  *DeviceComplianceDeviceStatusStatus `json:"status,omitempty"`
	UserName                                *string                             `json:"userName,omitempty"`
	UserPrincipalName                       *string                             `json:"userPrincipalName,omitempty"`
}
