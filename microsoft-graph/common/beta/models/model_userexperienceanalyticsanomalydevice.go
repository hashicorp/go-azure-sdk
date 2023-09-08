package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyDevice struct {
	AnomalyId                               *string                                           `json:"anomalyId,omitempty"`
	AnomalyOnDeviceFirstOccurrenceDateTime  *string                                           `json:"anomalyOnDeviceFirstOccurrenceDateTime,omitempty"`
	AnomalyOnDeviceLatestOccurrenceDateTime *string                                           `json:"anomalyOnDeviceLatestOccurrenceDateTime,omitempty"`
	CorrelationGroupId                      *string                                           `json:"correlationGroupId,omitempty"`
	DeviceId                                *string                                           `json:"deviceId,omitempty"`
	DeviceManufacturer                      *string                                           `json:"deviceManufacturer,omitempty"`
	DeviceModel                             *string                                           `json:"deviceModel,omitempty"`
	DeviceName                              *string                                           `json:"deviceName,omitempty"`
	DeviceStatus                            *UserExperienceAnalyticsAnomalyDeviceDeviceStatus `json:"deviceStatus,omitempty"`
	Id                                      *string                                           `json:"id,omitempty"`
	ODataType                               *string                                           `json:"@odata.type,omitempty"`
	OsName                                  *string                                           `json:"osName,omitempty"`
	OsVersion                               *string                                           `json:"osVersion,omitempty"`
}
