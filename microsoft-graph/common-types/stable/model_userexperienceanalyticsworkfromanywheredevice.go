package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereDevice struct {
	AutoPilotProfileAssigned      *bool                                                            `json:"autoPilotProfileAssigned,omitempty"`
	AutoPilotRegistered           *bool                                                            `json:"autoPilotRegistered,omitempty"`
	AzureAdDeviceId               *string                                                          `json:"azureAdDeviceId,omitempty"`
	AzureAdJoinType               *string                                                          `json:"azureAdJoinType,omitempty"`
	AzureAdRegistered             *bool                                                            `json:"azureAdRegistered,omitempty"`
	CloudIdentityScore            *float64                                                         `json:"cloudIdentityScore,omitempty"`
	CloudManagementScore          *float64                                                         `json:"cloudManagementScore,omitempty"`
	CloudProvisioningScore        *float64                                                         `json:"cloudProvisioningScore,omitempty"`
	CompliancePolicySetToIntune   *bool                                                            `json:"compliancePolicySetToIntune,omitempty"`
	DeviceId                      *string                                                          `json:"deviceId,omitempty"`
	DeviceName                    *string                                                          `json:"deviceName,omitempty"`
	HealthStatus                  *UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus       `json:"healthStatus,omitempty"`
	Id                            *string                                                          `json:"id,omitempty"`
	IsCloudManagedGatewayEnabled  *bool                                                            `json:"isCloudManagedGatewayEnabled,omitempty"`
	ManagedBy                     *string                                                          `json:"managedBy,omitempty"`
	Manufacturer                  *string                                                          `json:"manufacturer,omitempty"`
	Model                         *string                                                          `json:"model,omitempty"`
	ODataType                     *string                                                          `json:"@odata.type,omitempty"`
	OsCheckFailed                 *bool                                                            `json:"osCheckFailed,omitempty"`
	OsDescription                 *string                                                          `json:"osDescription,omitempty"`
	OsVersion                     *string                                                          `json:"osVersion,omitempty"`
	OtherWorkloadsSetToIntune     *bool                                                            `json:"otherWorkloadsSetToIntune,omitempty"`
	Ownership                     *string                                                          `json:"ownership,omitempty"`
	Processor64BitCheckFailed     *bool                                                            `json:"processor64BitCheckFailed,omitempty"`
	ProcessorCoreCountCheckFailed *bool                                                            `json:"processorCoreCountCheckFailed,omitempty"`
	ProcessorFamilyCheckFailed    *bool                                                            `json:"processorFamilyCheckFailed,omitempty"`
	ProcessorSpeedCheckFailed     *bool                                                            `json:"processorSpeedCheckFailed,omitempty"`
	RamCheckFailed                *bool                                                            `json:"ramCheckFailed,omitempty"`
	SecureBootCheckFailed         *bool                                                            `json:"secureBootCheckFailed,omitempty"`
	SerialNumber                  *string                                                          `json:"serialNumber,omitempty"`
	StorageCheckFailed            *bool                                                            `json:"storageCheckFailed,omitempty"`
	TenantAttached                *bool                                                            `json:"tenantAttached,omitempty"`
	TpmCheckFailed                *bool                                                            `json:"tpmCheckFailed,omitempty"`
	UpgradeEligibility            *UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility `json:"upgradeEligibility,omitempty"`
	WindowsScore                  *float64                                                         `json:"windowsScore,omitempty"`
	WorkFromAnywhereScore         *float64                                                         `json:"workFromAnywhereScore,omitempty"`
}
