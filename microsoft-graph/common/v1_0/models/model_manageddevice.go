package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDevice struct {
	ActivationLockBypassCode                  *string                                    `json:"activationLockBypassCode,omitempty"`
	AndroidSecurityPatchLevel                 *string                                    `json:"androidSecurityPatchLevel,omitempty"`
	AzureADDeviceId                           *string                                    `json:"azureADDeviceId,omitempty"`
	AzureADRegistered                         *bool                                      `json:"azureADRegistered,omitempty"`
	ComplianceGracePeriodExpirationDateTime   *string                                    `json:"complianceGracePeriodExpirationDateTime,omitempty"`
	ComplianceState                           *ManagedDeviceComplianceState              `json:"complianceState,omitempty"`
	ConfigurationManagerClientEnabledFeatures *ConfigurationManagerClientEnabledFeatures `json:"configurationManagerClientEnabledFeatures,omitempty"`
	DeviceActionResults                       *[]DeviceActionResult                      `json:"deviceActionResults,omitempty"`
	DeviceCategory                            *DeviceCategory                            `json:"deviceCategory,omitempty"`
	DeviceCategoryDisplayName                 *string                                    `json:"deviceCategoryDisplayName,omitempty"`
	DeviceCompliancePolicyStates              *[]DeviceCompliancePolicyState             `json:"deviceCompliancePolicyStates,omitempty"`
	DeviceConfigurationStates                 *[]DeviceConfigurationState                `json:"deviceConfigurationStates,omitempty"`
	DeviceEnrollmentType                      *ManagedDeviceDeviceEnrollmentType         `json:"deviceEnrollmentType,omitempty"`
	DeviceHealthAttestationState              *DeviceHealthAttestationState              `json:"deviceHealthAttestationState,omitempty"`
	DeviceName                                *string                                    `json:"deviceName,omitempty"`
	DeviceRegistrationState                   *ManagedDeviceDeviceRegistrationState      `json:"deviceRegistrationState,omitempty"`
	EasActivated                              *bool                                      `json:"easActivated,omitempty"`
	EasActivationDateTime                     *string                                    `json:"easActivationDateTime,omitempty"`
	EasDeviceId                               *string                                    `json:"easDeviceId,omitempty"`
	EmailAddress                              *string                                    `json:"emailAddress,omitempty"`
	EnrolledDateTime                          *string                                    `json:"enrolledDateTime,omitempty"`
	EthernetMacAddress                        *string                                    `json:"ethernetMacAddress,omitempty"`
	ExchangeAccessState                       *ManagedDeviceExchangeAccessState          `json:"exchangeAccessState,omitempty"`
	ExchangeAccessStateReason                 *ManagedDeviceExchangeAccessStateReason    `json:"exchangeAccessStateReason,omitempty"`
	ExchangeLastSuccessfulSyncDateTime        *string                                    `json:"exchangeLastSuccessfulSyncDateTime,omitempty"`
	FreeStorageSpaceInBytes                   *int64                                     `json:"freeStorageSpaceInBytes,omitempty"`
	Iccid                                     *string                                    `json:"iccid,omitempty"`
	Id                                        *string                                    `json:"id,omitempty"`
	Imei                                      *string                                    `json:"imei,omitempty"`
	IsEncrypted                               *bool                                      `json:"isEncrypted,omitempty"`
	IsSupervised                              *bool                                      `json:"isSupervised,omitempty"`
	JailBroken                                *string                                    `json:"jailBroken,omitempty"`
	LastSyncDateTime                          *string                                    `json:"lastSyncDateTime,omitempty"`
	LogCollectionRequests                     *[]DeviceLogCollectionResponse             `json:"logCollectionRequests,omitempty"`
	ManagedDeviceName                         *string                                    `json:"managedDeviceName,omitempty"`
	ManagedDeviceOwnerType                    *ManagedDeviceManagedDeviceOwnerType       `json:"managedDeviceOwnerType,omitempty"`
	ManagementAgent                           *ManagedDeviceManagementAgent              `json:"managementAgent,omitempty"`
	ManagementCertificateExpirationDate       *string                                    `json:"managementCertificateExpirationDate,omitempty"`
	Manufacturer                              *string                                    `json:"manufacturer,omitempty"`
	Meid                                      *string                                    `json:"meid,omitempty"`
	Model                                     *string                                    `json:"model,omitempty"`
	Notes                                     *string                                    `json:"notes,omitempty"`
	ODataType                                 *string                                    `json:"@odata.type,omitempty"`
	OperatingSystem                           *string                                    `json:"operatingSystem,omitempty"`
	OsVersion                                 *string                                    `json:"osVersion,omitempty"`
	PartnerReportedThreatState                *ManagedDevicePartnerReportedThreatState   `json:"partnerReportedThreatState,omitempty"`
	PhoneNumber                               *string                                    `json:"phoneNumber,omitempty"`
	PhysicalMemoryInBytes                     *int64                                     `json:"physicalMemoryInBytes,omitempty"`
	RemoteAssistanceSessionErrorDetails       *string                                    `json:"remoteAssistanceSessionErrorDetails,omitempty"`
	RemoteAssistanceSessionUrl                *string                                    `json:"remoteAssistanceSessionUrl,omitempty"`
	RequireUserEnrollmentApproval             *bool                                      `json:"requireUserEnrollmentApproval,omitempty"`
	SerialNumber                              *string                                    `json:"serialNumber,omitempty"`
	SubscriberCarrier                         *string                                    `json:"subscriberCarrier,omitempty"`
	TotalStorageSpaceInBytes                  *int64                                     `json:"totalStorageSpaceInBytes,omitempty"`
	Udid                                      *string                                    `json:"udid,omitempty"`
	UserDisplayName                           *string                                    `json:"userDisplayName,omitempty"`
	UserId                                    *string                                    `json:"userId,omitempty"`
	UserPrincipalName                         *string                                    `json:"userPrincipalName,omitempty"`
	Users                                     *[]User                                    `json:"users,omitempty"`
	WiFiMacAddress                            *string                                    `json:"wiFiMacAddress,omitempty"`
	WindowsProtectionState                    *WindowsProtectionState                    `json:"windowsProtectionState,omitempty"`
}
