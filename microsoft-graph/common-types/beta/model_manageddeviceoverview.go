package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceOverview struct {
	DeviceExchangeAccessStateSummary    *DeviceExchangeAccessStateSummary    `json:"deviceExchangeAccessStateSummary,omitempty"`
	DeviceOperatingSystemSummary        *DeviceOperatingSystemSummary        `json:"deviceOperatingSystemSummary,omitempty"`
	DualEnrolledDeviceCount             *int64                               `json:"dualEnrolledDeviceCount,omitempty"`
	EnrolledDeviceCount                 *int64                               `json:"enrolledDeviceCount,omitempty"`
	Id                                  *string                              `json:"id,omitempty"`
	LastModifiedDateTime                *string                              `json:"lastModifiedDateTime,omitempty"`
	ManagedDeviceModelsAndManufacturers *ManagedDeviceModelsAndManufacturers `json:"managedDeviceModelsAndManufacturers,omitempty"`
	MdmEnrolledCount                    *int64                               `json:"mdmEnrolledCount,omitempty"`
	ODataType                           *string                              `json:"@odata.type,omitempty"`
}
