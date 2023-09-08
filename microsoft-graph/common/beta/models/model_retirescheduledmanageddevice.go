package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetireScheduledManagedDevice struct {
	ComplianceState            *RetireScheduledManagedDeviceComplianceState `json:"complianceState,omitempty"`
	DeviceCompliancePolicyId   *string                                      `json:"deviceCompliancePolicyId,omitempty"`
	DeviceCompliancePolicyName *string                                      `json:"deviceCompliancePolicyName,omitempty"`
	DeviceType                 *RetireScheduledManagedDeviceDeviceType      `json:"deviceType,omitempty"`
	Id                         *string                                      `json:"id,omitempty"`
	ManagedDeviceId            *string                                      `json:"managedDeviceId,omitempty"`
	ManagedDeviceName          *string                                      `json:"managedDeviceName,omitempty"`
	ManagementAgent            *RetireScheduledManagedDeviceManagementAgent `json:"managementAgent,omitempty"`
	ODataType                  *string                                      `json:"@odata.type,omitempty"`
	OwnerType                  *RetireScheduledManagedDeviceOwnerType       `json:"ownerType,omitempty"`
	RetireAfterDateTime        *string                                      `json:"retireAfterDateTime,omitempty"`
	RoleScopeTagIds            *[]string                                    `json:"roleScopeTagIds,omitempty"`
}
