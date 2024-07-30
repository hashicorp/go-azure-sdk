package manageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindowsDeviceAccountActionParameter struct {
	CalendarSyncEnabled              *bool                 `json:"calendarSyncEnabled,omitempty"`
	DeviceAccount                    *WindowsDeviceAccount `json:"deviceAccount,omitempty"`
	DeviceAccountEmail               *string               `json:"deviceAccountEmail,omitempty"`
	ExchangeServer                   *string               `json:"exchangeServer,omitempty"`
	ODataType                        *string               `json:"@odata.type,omitempty"`
	PasswordRotationEnabled          *bool                 `json:"passwordRotationEnabled,omitempty"`
	SessionInitiationProtocalAddress *string               `json:"sessionInitiationProtocalAddress,omitempty"`
}
