package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceOperationPredicate struct {
	AadRegistered                               *bool
	ActivationLockBypassCode                    *string
	AndroidSecurityPatchLevel                   *string
	AutopilotEnrolled                           *bool
	AzureADDeviceId                             *string
	AzureADRegistered                           *bool
	AzureActiveDirectoryDeviceId                *string
	BootstrapTokenEscrowed                      *bool
	ComplianceGracePeriodExpirationDateTime     *string
	DeviceCategoryDisplayName                   *string
	DeviceFirmwareConfigurationInterfaceManaged *bool
	DeviceName                                  *string
	EasActivated                                *bool
	EasActivationDateTime                       *string
	EasDeviceId                                 *string
	EmailAddress                                *string
	EnrolledDateTime                            *string
	EnrollmentProfileName                       *string
	EthernetMacAddress                          *string
	ExchangeLastSuccessfulSyncDateTime          *string
	FreeStorageSpaceInBytes                     *int64
	Iccid                                       *string
	Id                                          *string
	Imei                                        *string
	IsEncrypted                                 *bool
	IsSupervised                                *bool
	JailBroken                                  *string
	LastSyncDateTime                            *string
	ManagedDeviceName                           *string
	ManagementCertificateExpirationDate         *string
	Manufacturer                                *string
	Meid                                        *string
	Model                                       *string
	Notes                                       *string
	ODataType                                   *string
	OperatingSystem                             *string
	OsVersion                                   *string
	PhoneNumber                                 *string
	PhysicalMemoryInBytes                       *int64
	PreferMdmOverGroupPolicyAppliedDateTime     *string
	RemoteAssistanceSessionErrorDetails         *string
	RemoteAssistanceSessionUrl                  *string
	RequireUserEnrollmentApproval               *bool
	RetireAfterDateTime                         *string
	SecurityPatchLevel                          *string
	SerialNumber                                *string
	SkuFamily                                   *string
	SkuNumber                                   *int64
	SpecificationVersion                        *string
	SubscriberCarrier                           *string
	TotalStorageSpaceInBytes                    *int64
	Udid                                        *string
	UserDisplayName                             *string
	UserId                                      *string
	UserPrincipalName                           *string
	WiFiMacAddress                              *string
	WindowsActiveMalwareCount                   *int64
	WindowsRemediatedMalwareCount               *int64
}

func (p ManagedDeviceOperationPredicate) Matches(input ManagedDevice) bool {

	if p.AadRegistered != nil && (input.AadRegistered == nil || *p.AadRegistered != *input.AadRegistered) {
		return false
	}

	if p.ActivationLockBypassCode != nil && (input.ActivationLockBypassCode == nil || *p.ActivationLockBypassCode != *input.ActivationLockBypassCode) {
		return false
	}

	if p.AndroidSecurityPatchLevel != nil && (input.AndroidSecurityPatchLevel == nil || *p.AndroidSecurityPatchLevel != *input.AndroidSecurityPatchLevel) {
		return false
	}

	if p.AutopilotEnrolled != nil && (input.AutopilotEnrolled == nil || *p.AutopilotEnrolled != *input.AutopilotEnrolled) {
		return false
	}

	if p.AzureADDeviceId != nil && (input.AzureADDeviceId == nil || *p.AzureADDeviceId != *input.AzureADDeviceId) {
		return false
	}

	if p.AzureADRegistered != nil && (input.AzureADRegistered == nil || *p.AzureADRegistered != *input.AzureADRegistered) {
		return false
	}

	if p.AzureActiveDirectoryDeviceId != nil && (input.AzureActiveDirectoryDeviceId == nil || *p.AzureActiveDirectoryDeviceId != *input.AzureActiveDirectoryDeviceId) {
		return false
	}

	if p.BootstrapTokenEscrowed != nil && (input.BootstrapTokenEscrowed == nil || *p.BootstrapTokenEscrowed != *input.BootstrapTokenEscrowed) {
		return false
	}

	if p.ComplianceGracePeriodExpirationDateTime != nil && (input.ComplianceGracePeriodExpirationDateTime == nil || *p.ComplianceGracePeriodExpirationDateTime != *input.ComplianceGracePeriodExpirationDateTime) {
		return false
	}

	if p.DeviceCategoryDisplayName != nil && (input.DeviceCategoryDisplayName == nil || *p.DeviceCategoryDisplayName != *input.DeviceCategoryDisplayName) {
		return false
	}

	if p.DeviceFirmwareConfigurationInterfaceManaged != nil && (input.DeviceFirmwareConfigurationInterfaceManaged == nil || *p.DeviceFirmwareConfigurationInterfaceManaged != *input.DeviceFirmwareConfigurationInterfaceManaged) {
		return false
	}

	if p.DeviceName != nil && (input.DeviceName == nil || *p.DeviceName != *input.DeviceName) {
		return false
	}

	if p.EasActivated != nil && (input.EasActivated == nil || *p.EasActivated != *input.EasActivated) {
		return false
	}

	if p.EasActivationDateTime != nil && (input.EasActivationDateTime == nil || *p.EasActivationDateTime != *input.EasActivationDateTime) {
		return false
	}

	if p.EasDeviceId != nil && (input.EasDeviceId == nil || *p.EasDeviceId != *input.EasDeviceId) {
		return false
	}

	if p.EmailAddress != nil && (input.EmailAddress == nil || *p.EmailAddress != *input.EmailAddress) {
		return false
	}

	if p.EnrolledDateTime != nil && (input.EnrolledDateTime == nil || *p.EnrolledDateTime != *input.EnrolledDateTime) {
		return false
	}

	if p.EnrollmentProfileName != nil && (input.EnrollmentProfileName == nil || *p.EnrollmentProfileName != *input.EnrollmentProfileName) {
		return false
	}

	if p.EthernetMacAddress != nil && (input.EthernetMacAddress == nil || *p.EthernetMacAddress != *input.EthernetMacAddress) {
		return false
	}

	if p.ExchangeLastSuccessfulSyncDateTime != nil && (input.ExchangeLastSuccessfulSyncDateTime == nil || *p.ExchangeLastSuccessfulSyncDateTime != *input.ExchangeLastSuccessfulSyncDateTime) {
		return false
	}

	if p.FreeStorageSpaceInBytes != nil && (input.FreeStorageSpaceInBytes == nil || *p.FreeStorageSpaceInBytes != *input.FreeStorageSpaceInBytes) {
		return false
	}

	if p.Iccid != nil && (input.Iccid == nil || *p.Iccid != *input.Iccid) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Imei != nil && (input.Imei == nil || *p.Imei != *input.Imei) {
		return false
	}

	if p.IsEncrypted != nil && (input.IsEncrypted == nil || *p.IsEncrypted != *input.IsEncrypted) {
		return false
	}

	if p.IsSupervised != nil && (input.IsSupervised == nil || *p.IsSupervised != *input.IsSupervised) {
		return false
	}

	if p.JailBroken != nil && (input.JailBroken == nil || *p.JailBroken != *input.JailBroken) {
		return false
	}

	if p.LastSyncDateTime != nil && (input.LastSyncDateTime == nil || *p.LastSyncDateTime != *input.LastSyncDateTime) {
		return false
	}

	if p.ManagedDeviceName != nil && (input.ManagedDeviceName == nil || *p.ManagedDeviceName != *input.ManagedDeviceName) {
		return false
	}

	if p.ManagementCertificateExpirationDate != nil && (input.ManagementCertificateExpirationDate == nil || *p.ManagementCertificateExpirationDate != *input.ManagementCertificateExpirationDate) {
		return false
	}

	if p.Manufacturer != nil && (input.Manufacturer == nil || *p.Manufacturer != *input.Manufacturer) {
		return false
	}

	if p.Meid != nil && (input.Meid == nil || *p.Meid != *input.Meid) {
		return false
	}

	if p.Model != nil && (input.Model == nil || *p.Model != *input.Model) {
		return false
	}

	if p.Notes != nil && (input.Notes == nil || *p.Notes != *input.Notes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OperatingSystem != nil && (input.OperatingSystem == nil || *p.OperatingSystem != *input.OperatingSystem) {
		return false
	}

	if p.OsVersion != nil && (input.OsVersion == nil || *p.OsVersion != *input.OsVersion) {
		return false
	}

	if p.PhoneNumber != nil && (input.PhoneNumber == nil || *p.PhoneNumber != *input.PhoneNumber) {
		return false
	}

	if p.PhysicalMemoryInBytes != nil && (input.PhysicalMemoryInBytes == nil || *p.PhysicalMemoryInBytes != *input.PhysicalMemoryInBytes) {
		return false
	}

	if p.PreferMdmOverGroupPolicyAppliedDateTime != nil && (input.PreferMdmOverGroupPolicyAppliedDateTime == nil || *p.PreferMdmOverGroupPolicyAppliedDateTime != *input.PreferMdmOverGroupPolicyAppliedDateTime) {
		return false
	}

	if p.RemoteAssistanceSessionErrorDetails != nil && (input.RemoteAssistanceSessionErrorDetails == nil || *p.RemoteAssistanceSessionErrorDetails != *input.RemoteAssistanceSessionErrorDetails) {
		return false
	}

	if p.RemoteAssistanceSessionUrl != nil && (input.RemoteAssistanceSessionUrl == nil || *p.RemoteAssistanceSessionUrl != *input.RemoteAssistanceSessionUrl) {
		return false
	}

	if p.RequireUserEnrollmentApproval != nil && (input.RequireUserEnrollmentApproval == nil || *p.RequireUserEnrollmentApproval != *input.RequireUserEnrollmentApproval) {
		return false
	}

	if p.RetireAfterDateTime != nil && (input.RetireAfterDateTime == nil || *p.RetireAfterDateTime != *input.RetireAfterDateTime) {
		return false
	}

	if p.SecurityPatchLevel != nil && (input.SecurityPatchLevel == nil || *p.SecurityPatchLevel != *input.SecurityPatchLevel) {
		return false
	}

	if p.SerialNumber != nil && (input.SerialNumber == nil || *p.SerialNumber != *input.SerialNumber) {
		return false
	}

	if p.SkuFamily != nil && (input.SkuFamily == nil || *p.SkuFamily != *input.SkuFamily) {
		return false
	}

	if p.SkuNumber != nil && (input.SkuNumber == nil || *p.SkuNumber != *input.SkuNumber) {
		return false
	}

	if p.SpecificationVersion != nil && (input.SpecificationVersion == nil || *p.SpecificationVersion != *input.SpecificationVersion) {
		return false
	}

	if p.SubscriberCarrier != nil && (input.SubscriberCarrier == nil || *p.SubscriberCarrier != *input.SubscriberCarrier) {
		return false
	}

	if p.TotalStorageSpaceInBytes != nil && (input.TotalStorageSpaceInBytes == nil || *p.TotalStorageSpaceInBytes != *input.TotalStorageSpaceInBytes) {
		return false
	}

	if p.Udid != nil && (input.Udid == nil || *p.Udid != *input.Udid) {
		return false
	}

	if p.UserDisplayName != nil && (input.UserDisplayName == nil || *p.UserDisplayName != *input.UserDisplayName) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	if p.WiFiMacAddress != nil && (input.WiFiMacAddress == nil || *p.WiFiMacAddress != *input.WiFiMacAddress) {
		return false
	}

	if p.WindowsActiveMalwareCount != nil && (input.WindowsActiveMalwareCount == nil || *p.WindowsActiveMalwareCount != *input.WindowsActiveMalwareCount) {
		return false
	}

	if p.WindowsRemediatedMalwareCount != nil && (input.WindowsRemediatedMalwareCount == nil || *p.WindowsRemediatedMalwareCount != *input.WindowsRemediatedMalwareCount) {
		return false
	}

	return true
}
