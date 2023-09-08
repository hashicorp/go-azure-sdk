package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeAccessRule struct {
	AccessLevel *DeviceManagementExchangeAccessRuleAccessLevel `json:"accessLevel,omitempty"`
	DeviceClass *DeviceManagementExchangeDeviceClass           `json:"deviceClass,omitempty"`
	ODataType   *string                                        `json:"@odata.type,omitempty"`
}
