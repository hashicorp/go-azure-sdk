package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceModelsAndManufacturers struct {
	DeviceManufacturers *[]string `json:"deviceManufacturers,omitempty"`
	DeviceModels        *[]string `json:"deviceModels,omitempty"`
	ODataType           *string   `json:"@odata.type,omitempty"`
}
