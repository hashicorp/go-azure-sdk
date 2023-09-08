package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsResourcePerformanceOperationPredicate struct {
	AverageSpikeTimeScore          *int64
	CpuSpikeTimeScore              *int64
	DeviceCount                    *int64
	DeviceId                       *string
	DeviceName                     *string
	DeviceResourcePerformanceScore *int64
	Id                             *string
	Manufacturer                   *string
	Model                          *string
	ODataType                      *string
	RamSpikeTimeScore              *int64
}

func (p UserExperienceAnalyticsResourcePerformanceOperationPredicate) Matches(input UserExperienceAnalyticsResourcePerformance) bool {

	if p.AverageSpikeTimeScore != nil && (input.AverageSpikeTimeScore == nil || *p.AverageSpikeTimeScore != *input.AverageSpikeTimeScore) {
		return false
	}

	if p.CpuSpikeTimeScore != nil && (input.CpuSpikeTimeScore == nil || *p.CpuSpikeTimeScore != *input.CpuSpikeTimeScore) {
		return false
	}

	if p.DeviceCount != nil && (input.DeviceCount == nil || *p.DeviceCount != *input.DeviceCount) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.DeviceName != nil && (input.DeviceName == nil || *p.DeviceName != *input.DeviceName) {
		return false
	}

	if p.DeviceResourcePerformanceScore != nil && (input.DeviceResourcePerformanceScore == nil || *p.DeviceResourcePerformanceScore != *input.DeviceResourcePerformanceScore) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Manufacturer != nil && (input.Manufacturer == nil || *p.Manufacturer != *input.Manufacturer) {
		return false
	}

	if p.Model != nil && (input.Model == nil || *p.Model != *input.Model) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RamSpikeTimeScore != nil && (input.RamSpikeTimeScore == nil || *p.RamSpikeTimeScore != *input.RamSpikeTimeScore) {
		return false
	}

	return true
}
