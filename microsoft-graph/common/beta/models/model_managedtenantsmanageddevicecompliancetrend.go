package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedDeviceComplianceTrend struct {
	CompliantDeviceCount     *int64  `json:"compliantDeviceCount,omitempty"`
	ConfigManagerDeviceCount *int64  `json:"configManagerDeviceCount,omitempty"`
	CountDateTime            *string `json:"countDateTime,omitempty"`
	ErrorDeviceCount         *int64  `json:"errorDeviceCount,omitempty"`
	Id                       *string `json:"id,omitempty"`
	InGracePeriodDeviceCount *int64  `json:"inGracePeriodDeviceCount,omitempty"`
	NoncompliantDeviceCount  *int64  `json:"noncompliantDeviceCount,omitempty"`
	ODataType                *string `json:"@odata.type,omitempty"`
	TenantDisplayName        *string `json:"tenantDisplayName,omitempty"`
	TenantId                 *string `json:"tenantId,omitempty"`
	UnknownDeviceCount       *int64  `json:"unknownDeviceCount,omitempty"`
}
