package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthModelPerformanceOperationPredicate struct {
	ActiveDevices                    *int64
	AverageBatteryAgeInDays          *int64
	AverageEstimatedRuntimeInMinutes *int64
	AverageMaxCapacityPercentage     *int64
	Id                               *string
	Manufacturer                     *string
	MeanFullBatteryDrainCount        *int64
	MedianEstimatedRuntimeInMinutes  *int64
	MedianFullBatteryDrainCount      *int64
	MedianMaxCapacityPercentage      *int64
	Model                            *string
	ModelBatteryHealthScore          *int64
	ODataType                        *string
}

func (p UserExperienceAnalyticsBatteryHealthModelPerformanceOperationPredicate) Matches(input UserExperienceAnalyticsBatteryHealthModelPerformance) bool {

	if p.ActiveDevices != nil && (input.ActiveDevices == nil || *p.ActiveDevices != *input.ActiveDevices) {
		return false
	}

	if p.AverageBatteryAgeInDays != nil && (input.AverageBatteryAgeInDays == nil || *p.AverageBatteryAgeInDays != *input.AverageBatteryAgeInDays) {
		return false
	}

	if p.AverageEstimatedRuntimeInMinutes != nil && (input.AverageEstimatedRuntimeInMinutes == nil || *p.AverageEstimatedRuntimeInMinutes != *input.AverageEstimatedRuntimeInMinutes) {
		return false
	}

	if p.AverageMaxCapacityPercentage != nil && (input.AverageMaxCapacityPercentage == nil || *p.AverageMaxCapacityPercentage != *input.AverageMaxCapacityPercentage) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Manufacturer != nil && (input.Manufacturer == nil || *p.Manufacturer != *input.Manufacturer) {
		return false
	}

	if p.MeanFullBatteryDrainCount != nil && (input.MeanFullBatteryDrainCount == nil || *p.MeanFullBatteryDrainCount != *input.MeanFullBatteryDrainCount) {
		return false
	}

	if p.MedianEstimatedRuntimeInMinutes != nil && (input.MedianEstimatedRuntimeInMinutes == nil || *p.MedianEstimatedRuntimeInMinutes != *input.MedianEstimatedRuntimeInMinutes) {
		return false
	}

	if p.MedianFullBatteryDrainCount != nil && (input.MedianFullBatteryDrainCount == nil || *p.MedianFullBatteryDrainCount != *input.MedianFullBatteryDrainCount) {
		return false
	}

	if p.MedianMaxCapacityPercentage != nil && (input.MedianMaxCapacityPercentage == nil || *p.MedianMaxCapacityPercentage != *input.MedianMaxCapacityPercentage) {
		return false
	}

	if p.Model != nil && (input.Model == nil || *p.Model != *input.Model) {
		return false
	}

	if p.ModelBatteryHealthScore != nil && (input.ModelBatteryHealthScore == nil || *p.ModelBatteryHealthScore != *input.ModelBatteryHealthScore) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
