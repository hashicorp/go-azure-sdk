package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentity struct {
	AddressableUserName                       *string                                                                  `json:"addressableUserName,omitempty"`
	AzureActiveDirectoryDeviceId              *string                                                                  `json:"azureActiveDirectoryDeviceId,omitempty"`
	AzureAdDeviceId                           *string                                                                  `json:"azureAdDeviceId,omitempty"`
	DeploymentProfile                         *WindowsAutopilotDeploymentProfile                                       `json:"deploymentProfile,omitempty"`
	DeploymentProfileAssignedDateTime         *string                                                                  `json:"deploymentProfileAssignedDateTime,omitempty"`
	DeploymentProfileAssignmentDetailedStatus *WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus `json:"deploymentProfileAssignmentDetailedStatus,omitempty"`
	DeploymentProfileAssignmentStatus         *WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus         `json:"deploymentProfileAssignmentStatus,omitempty"`
	DeviceAccountPassword                     *string                                                                  `json:"deviceAccountPassword,omitempty"`
	DeviceAccountUpn                          *string                                                                  `json:"deviceAccountUpn,omitempty"`
	DeviceFriendlyName                        *string                                                                  `json:"deviceFriendlyName,omitempty"`
	DisplayName                               *string                                                                  `json:"displayName,omitempty"`
	EnrollmentState                           *WindowsAutopilotDeviceIdentityEnrollmentState                           `json:"enrollmentState,omitempty"`
	GroupTag                                  *string                                                                  `json:"groupTag,omitempty"`
	Id                                        *string                                                                  `json:"id,omitempty"`
	IntendedDeploymentProfile                 *WindowsAutopilotDeploymentProfile                                       `json:"intendedDeploymentProfile,omitempty"`
	LastContactedDateTime                     *string                                                                  `json:"lastContactedDateTime,omitempty"`
	ManagedDeviceId                           *string                                                                  `json:"managedDeviceId,omitempty"`
	Manufacturer                              *string                                                                  `json:"manufacturer,omitempty"`
	Model                                     *string                                                                  `json:"model,omitempty"`
	ODataType                                 *string                                                                  `json:"@odata.type,omitempty"`
	ProductKey                                *string                                                                  `json:"productKey,omitempty"`
	PurchaseOrderIdentifier                   *string                                                                  `json:"purchaseOrderIdentifier,omitempty"`
	RemediationState                          *WindowsAutopilotDeviceIdentityRemediationState                          `json:"remediationState,omitempty"`
	RemediationStateLastModifiedDateTime      *string                                                                  `json:"remediationStateLastModifiedDateTime,omitempty"`
	ResourceName                              *string                                                                  `json:"resourceName,omitempty"`
	SerialNumber                              *string                                                                  `json:"serialNumber,omitempty"`
	SkuNumber                                 *string                                                                  `json:"skuNumber,omitempty"`
	SystemFamily                              *string                                                                  `json:"systemFamily,omitempty"`
	UserPrincipalName                         *string                                                                  `json:"userPrincipalName,omitempty"`
	UserlessEnrollmentStatus                  *WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus                  `json:"userlessEnrollmentStatus,omitempty"`
}
