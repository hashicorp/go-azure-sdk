package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPC struct {
	AadDeviceId                *string                             `json:"aadDeviceId,omitempty"`
	ConnectionSettings         *CloudPCConnectionSettings          `json:"connectionSettings,omitempty"`
	ConnectivityResult         *CloudPCConnectivityResult          `json:"connectivityResult,omitempty"`
	DiskEncryptionState        *CloudPCDiskEncryptionState         `json:"diskEncryptionState,omitempty"`
	DisplayName                *string                             `json:"displayName,omitempty"`
	GracePeriodEndDateTime     *string                             `json:"gracePeriodEndDateTime,omitempty"`
	Id                         *string                             `json:"id,omitempty"`
	ImageDisplayName           *string                             `json:"imageDisplayName,omitempty"`
	LastLoginResult            *CloudPCLoginResult                 `json:"lastLoginResult,omitempty"`
	LastModifiedDateTime       *string                             `json:"lastModifiedDateTime,omitempty"`
	LastRemoteActionResult     *CloudPCRemoteActionResult          `json:"lastRemoteActionResult,omitempty"`
	ManagedDeviceId            *string                             `json:"managedDeviceId,omitempty"`
	ManagedDeviceName          *string                             `json:"managedDeviceName,omitempty"`
	ODataType                  *string                             `json:"@odata.type,omitempty"`
	OnPremisesConnectionName   *string                             `json:"onPremisesConnectionName,omitempty"`
	OsVersion                  *CloudPCOsVersion                   `json:"osVersion,omitempty"`
	PartnerAgentInstallResults *[]CloudPCPartnerAgentInstallResult `json:"partnerAgentInstallResults,omitempty"`
	PowerState                 *CloudPCPowerState                  `json:"powerState,omitempty"`
	ProvisioningPolicyId       *string                             `json:"provisioningPolicyId,omitempty"`
	ProvisioningPolicyName     *string                             `json:"provisioningPolicyName,omitempty"`
	ProvisioningType           *CloudPCProvisioningType            `json:"provisioningType,omitempty"`
	ScopeIds                   *[]string                           `json:"scopeIds,omitempty"`
	ServicePlanId              *string                             `json:"servicePlanId,omitempty"`
	ServicePlanName            *string                             `json:"servicePlanName,omitempty"`
	ServicePlanType            *CloudPCServicePlanType             `json:"servicePlanType,omitempty"`
	Status                     *CloudPCStatus                      `json:"status,omitempty"`
	StatusDetails              *CloudPCStatusDetails               `json:"statusDetails,omitempty"`
	UserAccountType            *CloudPCUserAccountType             `json:"userAccountType,omitempty"`
	UserPrincipalName          *string                             `json:"userPrincipalName,omitempty"`
}
