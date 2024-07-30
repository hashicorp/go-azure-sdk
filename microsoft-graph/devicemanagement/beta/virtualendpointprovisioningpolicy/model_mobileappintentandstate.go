package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppIntentAndState struct {
	Id                      *string                          `json:"id,omitempty"`
	ManagedDeviceIdentifier *string                          `json:"managedDeviceIdentifier,omitempty"`
	MobileAppList           *[]MobileAppIntentAndStateDetail `json:"mobileAppList,omitempty"`
	ODataType               *string                          `json:"@odata.type,omitempty"`
	UserId                  *string                          `json:"userId,omitempty"`
}
