package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeDeviceClass struct {
	Name      *string                                  `json:"name,omitempty"`
	ODataType *string                                  `json:"@odata.type,omitempty"`
	Type      *DeviceManagementExchangeDeviceClassType `json:"type,omitempty"`
}
