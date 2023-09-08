package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationBandwidthBusinessHoursLimit struct {
	BandwidthBeginBusinessHours             *int64  `json:"bandwidthBeginBusinessHours,omitempty"`
	BandwidthEndBusinessHours               *int64  `json:"bandwidthEndBusinessHours,omitempty"`
	BandwidthPercentageDuringBusinessHours  *int64  `json:"bandwidthPercentageDuringBusinessHours,omitempty"`
	BandwidthPercentageOutsideBusinessHours *int64  `json:"bandwidthPercentageOutsideBusinessHours,omitempty"`
	ODataType                               *string `json:"@odata.type,omitempty"`
}
