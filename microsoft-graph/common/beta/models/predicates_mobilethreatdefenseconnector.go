package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileThreatDefenseConnectorOperationPredicate struct {
	AllowPartnerToCollectIOSApplicationMetadata         *bool
	AllowPartnerToCollectIOSPersonalApplicationMetadata *bool
	AndroidDeviceBlockedOnMissingPartnerData            *bool
	AndroidEnabled                                      *bool
	AndroidMobileApplicationManagementEnabled           *bool
	Id                                                  *string
	IosDeviceBlockedOnMissingPartnerData                *bool
	IosEnabled                                          *bool
	IosMobileApplicationManagementEnabled               *bool
	LastHeartbeatDateTime                               *string
	MacDeviceBlockedOnMissingPartnerData                *bool
	MacEnabled                                          *bool
	MicrosoftDefenderForEndpointAttachEnabled           *bool
	ODataType                                           *string
	PartnerUnresponsivenessThresholdInDays              *int64
	PartnerUnsupportedOsVersionBlocked                  *bool
	WindowsDeviceBlockedOnMissingPartnerData            *bool
	WindowsEnabled                                      *bool
	WindowsMobileApplicationManagementEnabled           *bool
}

func (p MobileThreatDefenseConnectorOperationPredicate) Matches(input MobileThreatDefenseConnector) bool {

	if p.AllowPartnerToCollectIOSApplicationMetadata != nil && (input.AllowPartnerToCollectIOSApplicationMetadata == nil || *p.AllowPartnerToCollectIOSApplicationMetadata != *input.AllowPartnerToCollectIOSApplicationMetadata) {
		return false
	}

	if p.AllowPartnerToCollectIOSPersonalApplicationMetadata != nil && (input.AllowPartnerToCollectIOSPersonalApplicationMetadata == nil || *p.AllowPartnerToCollectIOSPersonalApplicationMetadata != *input.AllowPartnerToCollectIOSPersonalApplicationMetadata) {
		return false
	}

	if p.AndroidDeviceBlockedOnMissingPartnerData != nil && (input.AndroidDeviceBlockedOnMissingPartnerData == nil || *p.AndroidDeviceBlockedOnMissingPartnerData != *input.AndroidDeviceBlockedOnMissingPartnerData) {
		return false
	}

	if p.AndroidEnabled != nil && (input.AndroidEnabled == nil || *p.AndroidEnabled != *input.AndroidEnabled) {
		return false
	}

	if p.AndroidMobileApplicationManagementEnabled != nil && (input.AndroidMobileApplicationManagementEnabled == nil || *p.AndroidMobileApplicationManagementEnabled != *input.AndroidMobileApplicationManagementEnabled) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IosDeviceBlockedOnMissingPartnerData != nil && (input.IosDeviceBlockedOnMissingPartnerData == nil || *p.IosDeviceBlockedOnMissingPartnerData != *input.IosDeviceBlockedOnMissingPartnerData) {
		return false
	}

	if p.IosEnabled != nil && (input.IosEnabled == nil || *p.IosEnabled != *input.IosEnabled) {
		return false
	}

	if p.IosMobileApplicationManagementEnabled != nil && (input.IosMobileApplicationManagementEnabled == nil || *p.IosMobileApplicationManagementEnabled != *input.IosMobileApplicationManagementEnabled) {
		return false
	}

	if p.LastHeartbeatDateTime != nil && (input.LastHeartbeatDateTime == nil || *p.LastHeartbeatDateTime != *input.LastHeartbeatDateTime) {
		return false
	}

	if p.MacDeviceBlockedOnMissingPartnerData != nil && (input.MacDeviceBlockedOnMissingPartnerData == nil || *p.MacDeviceBlockedOnMissingPartnerData != *input.MacDeviceBlockedOnMissingPartnerData) {
		return false
	}

	if p.MacEnabled != nil && (input.MacEnabled == nil || *p.MacEnabled != *input.MacEnabled) {
		return false
	}

	if p.MicrosoftDefenderForEndpointAttachEnabled != nil && (input.MicrosoftDefenderForEndpointAttachEnabled == nil || *p.MicrosoftDefenderForEndpointAttachEnabled != *input.MicrosoftDefenderForEndpointAttachEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PartnerUnresponsivenessThresholdInDays != nil && (input.PartnerUnresponsivenessThresholdInDays == nil || *p.PartnerUnresponsivenessThresholdInDays != *input.PartnerUnresponsivenessThresholdInDays) {
		return false
	}

	if p.PartnerUnsupportedOsVersionBlocked != nil && (input.PartnerUnsupportedOsVersionBlocked == nil || *p.PartnerUnsupportedOsVersionBlocked != *input.PartnerUnsupportedOsVersionBlocked) {
		return false
	}

	if p.WindowsDeviceBlockedOnMissingPartnerData != nil && (input.WindowsDeviceBlockedOnMissingPartnerData == nil || *p.WindowsDeviceBlockedOnMissingPartnerData != *input.WindowsDeviceBlockedOnMissingPartnerData) {
		return false
	}

	if p.WindowsEnabled != nil && (input.WindowsEnabled == nil || *p.WindowsEnabled != *input.WindowsEnabled) {
		return false
	}

	if p.WindowsMobileApplicationManagementEnabled != nil && (input.WindowsMobileApplicationManagementEnabled == nil || *p.WindowsMobileApplicationManagementEnabled != *input.WindowsMobileApplicationManagementEnabled) {
		return false
	}

	return true
}
