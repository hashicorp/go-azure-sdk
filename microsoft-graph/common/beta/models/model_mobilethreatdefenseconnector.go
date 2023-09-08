package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileThreatDefenseConnector struct {
	AllowPartnerToCollectIOSApplicationMetadata         *bool                                     `json:"allowPartnerToCollectIOSApplicationMetadata,omitempty"`
	AllowPartnerToCollectIOSPersonalApplicationMetadata *bool                                     `json:"allowPartnerToCollectIOSPersonalApplicationMetadata,omitempty"`
	AndroidDeviceBlockedOnMissingPartnerData            *bool                                     `json:"androidDeviceBlockedOnMissingPartnerData,omitempty"`
	AndroidEnabled                                      *bool                                     `json:"androidEnabled,omitempty"`
	AndroidMobileApplicationManagementEnabled           *bool                                     `json:"androidMobileApplicationManagementEnabled,omitempty"`
	Id                                                  *string                                   `json:"id,omitempty"`
	IosDeviceBlockedOnMissingPartnerData                *bool                                     `json:"iosDeviceBlockedOnMissingPartnerData,omitempty"`
	IosEnabled                                          *bool                                     `json:"iosEnabled,omitempty"`
	IosMobileApplicationManagementEnabled               *bool                                     `json:"iosMobileApplicationManagementEnabled,omitempty"`
	LastHeartbeatDateTime                               *string                                   `json:"lastHeartbeatDateTime,omitempty"`
	MacDeviceBlockedOnMissingPartnerData                *bool                                     `json:"macDeviceBlockedOnMissingPartnerData,omitempty"`
	MacEnabled                                          *bool                                     `json:"macEnabled,omitempty"`
	MicrosoftDefenderForEndpointAttachEnabled           *bool                                     `json:"microsoftDefenderForEndpointAttachEnabled,omitempty"`
	ODataType                                           *string                                   `json:"@odata.type,omitempty"`
	PartnerState                                        *MobileThreatDefenseConnectorPartnerState `json:"partnerState,omitempty"`
	PartnerUnresponsivenessThresholdInDays              *int64                                    `json:"partnerUnresponsivenessThresholdInDays,omitempty"`
	PartnerUnsupportedOsVersionBlocked                  *bool                                     `json:"partnerUnsupportedOsVersionBlocked,omitempty"`
	WindowsDeviceBlockedOnMissingPartnerData            *bool                                     `json:"windowsDeviceBlockedOnMissingPartnerData,omitempty"`
	WindowsEnabled                                      *bool                                     `json:"windowsEnabled,omitempty"`
	WindowsMobileApplicationManagementEnabled           *bool                                     `json:"windowsMobileApplicationManagementEnabled,omitempty"`
}
