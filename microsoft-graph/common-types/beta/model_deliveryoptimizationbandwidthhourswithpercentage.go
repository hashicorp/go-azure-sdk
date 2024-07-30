package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationBandwidthHoursWithPercentage struct {
	BandwidthBackgroundPercentageHours *DeliveryOptimizationBandwidthBusinessHoursLimit `json:"bandwidthBackgroundPercentageHours,omitempty"`
	BandwidthForegroundPercentageHours *DeliveryOptimizationBandwidthBusinessHoursLimit `json:"bandwidthForegroundPercentageHours,omitempty"`
	ODataType                          *string                                          `json:"@odata.type,omitempty"`
}
