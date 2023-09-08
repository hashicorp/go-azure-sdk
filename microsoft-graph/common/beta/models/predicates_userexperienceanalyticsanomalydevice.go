package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyDeviceOperationPredicate struct {
	AnomalyId                               *string
	AnomalyOnDeviceFirstOccurrenceDateTime  *string
	AnomalyOnDeviceLatestOccurrenceDateTime *string
	CorrelationGroupId                      *string
	DeviceId                                *string
	DeviceManufacturer                      *string
	DeviceModel                             *string
	DeviceName                              *string
	Id                                      *string
	ODataType                               *string
	OsName                                  *string
	OsVersion                               *string
}

func (p UserExperienceAnalyticsAnomalyDeviceOperationPredicate) Matches(input UserExperienceAnalyticsAnomalyDevice) bool {

	if p.AnomalyId != nil && (input.AnomalyId == nil || *p.AnomalyId != *input.AnomalyId) {
		return false
	}

	if p.AnomalyOnDeviceFirstOccurrenceDateTime != nil && (input.AnomalyOnDeviceFirstOccurrenceDateTime == nil || *p.AnomalyOnDeviceFirstOccurrenceDateTime != *input.AnomalyOnDeviceFirstOccurrenceDateTime) {
		return false
	}

	if p.AnomalyOnDeviceLatestOccurrenceDateTime != nil && (input.AnomalyOnDeviceLatestOccurrenceDateTime == nil || *p.AnomalyOnDeviceLatestOccurrenceDateTime != *input.AnomalyOnDeviceLatestOccurrenceDateTime) {
		return false
	}

	if p.CorrelationGroupId != nil && (input.CorrelationGroupId == nil || *p.CorrelationGroupId != *input.CorrelationGroupId) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.DeviceManufacturer != nil && (input.DeviceManufacturer == nil || *p.DeviceManufacturer != *input.DeviceManufacturer) {
		return false
	}

	if p.DeviceModel != nil && (input.DeviceModel == nil || *p.DeviceModel != *input.DeviceModel) {
		return false
	}

	if p.DeviceName != nil && (input.DeviceName == nil || *p.DeviceName != *input.DeviceName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OsName != nil && (input.OsName == nil || *p.OsName != *input.OsName) {
		return false
	}

	if p.OsVersion != nil && (input.OsVersion == nil || *p.OsVersion != *input.OsVersion) {
		return false
	}

	return true
}
