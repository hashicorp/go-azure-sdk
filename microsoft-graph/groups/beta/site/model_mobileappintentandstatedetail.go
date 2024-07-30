package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppIntentAndStateDetail struct {
	ApplicationId        *string                                       `json:"applicationId,omitempty"`
	DisplayName          *string                                       `json:"displayName,omitempty"`
	DisplayVersion       *string                                       `json:"displayVersion,omitempty"`
	InstallState         *MobileAppIntentAndStateDetailInstallState    `json:"installState,omitempty"`
	MobileAppIntent      *MobileAppIntentAndStateDetailMobileAppIntent `json:"mobileAppIntent,omitempty"`
	ODataType            *string                                       `json:"@odata.type,omitempty"`
	SupportedDeviceTypes *[]MobileAppSupportedDeviceType               `json:"supportedDeviceTypes,omitempty"`
}
