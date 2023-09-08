package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationBandwidthPercentageOperationPredicate struct {
	MaximumBackgroundBandwidthPercentage *int64
	MaximumForegroundBandwidthPercentage *int64
	ODataType                            *string
}

func (p DeliveryOptimizationBandwidthPercentageOperationPredicate) Matches(input DeliveryOptimizationBandwidthPercentage) bool {

	if p.MaximumBackgroundBandwidthPercentage != nil && (input.MaximumBackgroundBandwidthPercentage == nil || *p.MaximumBackgroundBandwidthPercentage != *input.MaximumBackgroundBandwidthPercentage) {
		return false
	}

	if p.MaximumForegroundBandwidthPercentage != nil && (input.MaximumForegroundBandwidthPercentage == nil || *p.MaximumForegroundBandwidthPercentage != *input.MaximumForegroundBandwidthPercentage) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
