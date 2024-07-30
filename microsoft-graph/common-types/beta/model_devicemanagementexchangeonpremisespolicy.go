package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeOnPremisesPolicy struct {
	AccessRules               *[]DeviceManagementExchangeAccessRule                       `json:"accessRules,omitempty"`
	ConditionalAccessSettings *OnPremisesConditionalAccessSettings                        `json:"conditionalAccessSettings,omitempty"`
	DefaultAccessLevel        *DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel `json:"defaultAccessLevel,omitempty"`
	Id                        *string                                                     `json:"id,omitempty"`
	KnownDeviceClasses        *[]DeviceManagementExchangeDeviceClass                      `json:"knownDeviceClasses,omitempty"`
	NotificationContent       *string                                                     `json:"notificationContent,omitempty"`
	ODataType                 *string                                                     `json:"@odata.type,omitempty"`
}
