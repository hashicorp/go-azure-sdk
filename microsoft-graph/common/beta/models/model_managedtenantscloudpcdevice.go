package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsCloudPCDevice struct {
	CloudPCStatus         *string `json:"cloudPcStatus,omitempty"`
	DeviceSpecification   *string `json:"deviceSpecification,omitempty"`
	DisplayName           *string `json:"displayName,omitempty"`
	Id                    *string `json:"id,omitempty"`
	LastRefreshedDateTime *string `json:"lastRefreshedDateTime,omitempty"`
	ManagedDeviceId       *string `json:"managedDeviceId,omitempty"`
	ManagedDeviceName     *string `json:"managedDeviceName,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	ProvisioningPolicyId  *string `json:"provisioningPolicyId,omitempty"`
	ServicePlanName       *string `json:"servicePlanName,omitempty"`
	ServicePlanType       *string `json:"servicePlanType,omitempty"`
	TenantDisplayName     *string `json:"tenantDisplayName,omitempty"`
	TenantId              *string `json:"tenantId,omitempty"`
	UserPrincipalName     *string `json:"userPrincipalName,omitempty"`
}
