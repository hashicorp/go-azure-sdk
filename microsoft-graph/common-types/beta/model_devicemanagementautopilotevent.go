package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotEvent struct {
	AccountSetupDuration                                      *string                                           `json:"accountSetupDuration,omitempty"`
	AccountSetupStatus                                        *DeviceManagementAutopilotEventAccountSetupStatus `json:"accountSetupStatus,omitempty"`
	DeploymentDuration                                        *string                                           `json:"deploymentDuration,omitempty"`
	DeploymentEndDateTime                                     *string                                           `json:"deploymentEndDateTime,omitempty"`
	DeploymentStartDateTime                                   *string                                           `json:"deploymentStartDateTime,omitempty"`
	DeploymentState                                           *DeviceManagementAutopilotEventDeploymentState    `json:"deploymentState,omitempty"`
	DeploymentTotalDuration                                   *string                                           `json:"deploymentTotalDuration,omitempty"`
	DeviceId                                                  *string                                           `json:"deviceId,omitempty"`
	DevicePreparationDuration                                 *string                                           `json:"devicePreparationDuration,omitempty"`
	DeviceRegisteredDateTime                                  *string                                           `json:"deviceRegisteredDateTime,omitempty"`
	DeviceSerialNumber                                        *string                                           `json:"deviceSerialNumber,omitempty"`
	DeviceSetupDuration                                       *string                                           `json:"deviceSetupDuration,omitempty"`
	DeviceSetupStatus                                         *DeviceManagementAutopilotEventDeviceSetupStatus  `json:"deviceSetupStatus,omitempty"`
	EnrollmentFailureDetails                                  *string                                           `json:"enrollmentFailureDetails,omitempty"`
	EnrollmentStartDateTime                                   *string                                           `json:"enrollmentStartDateTime,omitempty"`
	EnrollmentState                                           *DeviceManagementAutopilotEventEnrollmentState    `json:"enrollmentState,omitempty"`
	EnrollmentType                                            *DeviceManagementAutopilotEventEnrollmentType     `json:"enrollmentType,omitempty"`
	EventDateTime                                             *string                                           `json:"eventDateTime,omitempty"`
	Id                                                        *string                                           `json:"id,omitempty"`
	ManagedDeviceName                                         *string                                           `json:"managedDeviceName,omitempty"`
	ODataType                                                 *string                                           `json:"@odata.type,omitempty"`
	OsVersion                                                 *string                                           `json:"osVersion,omitempty"`
	PolicyStatusDetails                                       *[]DeviceManagementAutopilotPolicyStatusDetail    `json:"policyStatusDetails,omitempty"`
	TargetedAppCount                                          *int64                                            `json:"targetedAppCount,omitempty"`
	TargetedPolicyCount                                       *int64                                            `json:"targetedPolicyCount,omitempty"`
	UserPrincipalName                                         *string                                           `json:"userPrincipalName,omitempty"`
	Windows10EnrollmentCompletionPageConfigurationDisplayName *string                                           `json:"windows10EnrollmentCompletionPageConfigurationDisplayName,omitempty"`
	Windows10EnrollmentCompletionPageConfigurationId          *string                                           `json:"windows10EnrollmentCompletionPageConfigurationId,omitempty"`
	WindowsAutopilotDeploymentProfileDisplayName              *string                                           `json:"windowsAutopilotDeploymentProfileDisplayName,omitempty"`
}
