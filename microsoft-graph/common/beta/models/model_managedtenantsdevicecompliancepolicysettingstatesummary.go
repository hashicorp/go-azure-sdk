package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsDeviceCompliancePolicySettingStateSummary struct {
	ConflictDeviceCount      *int64  `json:"conflictDeviceCount,omitempty"`
	ErrorDeviceCount         *int64  `json:"errorDeviceCount,omitempty"`
	FailedDeviceCount        *int64  `json:"failedDeviceCount,omitempty"`
	Id                       *string `json:"id,omitempty"`
	IntuneAccountId          *string `json:"intuneAccountId,omitempty"`
	IntuneSettingId          *string `json:"intuneSettingId,omitempty"`
	LastRefreshedDateTime    *string `json:"lastRefreshedDateTime,omitempty"`
	NotApplicableDeviceCount *int64  `json:"notApplicableDeviceCount,omitempty"`
	ODataType                *string `json:"@odata.type,omitempty"`
	PendingDeviceCount       *int64  `json:"pendingDeviceCount,omitempty"`
	PolicyType               *string `json:"policyType,omitempty"`
	SettingName              *string `json:"settingName,omitempty"`
	SucceededDeviceCount     *int64  `json:"succeededDeviceCount,omitempty"`
	TenantDisplayName        *string `json:"tenantDisplayName,omitempty"`
	TenantId                 *string `json:"tenantId,omitempty"`
}
