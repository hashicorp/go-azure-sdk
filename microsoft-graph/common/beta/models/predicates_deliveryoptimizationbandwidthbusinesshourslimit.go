package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationBandwidthBusinessHoursLimitOperationPredicate struct {
	BandwidthBeginBusinessHours             *int64
	BandwidthEndBusinessHours               *int64
	BandwidthPercentageDuringBusinessHours  *int64
	BandwidthPercentageOutsideBusinessHours *int64
	ODataType                               *string
}

func (p DeliveryOptimizationBandwidthBusinessHoursLimitOperationPredicate) Matches(input DeliveryOptimizationBandwidthBusinessHoursLimit) bool {

	if p.BandwidthBeginBusinessHours != nil && (input.BandwidthBeginBusinessHours == nil || *p.BandwidthBeginBusinessHours != *input.BandwidthBeginBusinessHours) {
		return false
	}

	if p.BandwidthEndBusinessHours != nil && (input.BandwidthEndBusinessHours == nil || *p.BandwidthEndBusinessHours != *input.BandwidthEndBusinessHours) {
		return false
	}

	if p.BandwidthPercentageDuringBusinessHours != nil && (input.BandwidthPercentageDuringBusinessHours == nil || *p.BandwidthPercentageDuringBusinessHours != *input.BandwidthPercentageDuringBusinessHours) {
		return false
	}

	if p.BandwidthPercentageOutsideBusinessHours != nil && (input.BandwidthPercentageOutsideBusinessHours == nil || *p.BandwidthPercentageOutsideBusinessHours != *input.BandwidthPercentageOutsideBusinessHours) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
