package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationBandwidthPercentage struct {
	MaximumBackgroundBandwidthPercentage *int64  `json:"maximumBackgroundBandwidthPercentage,omitempty"`
	MaximumForegroundBandwidthPercentage *int64  `json:"maximumForegroundBandwidthPercentage,omitempty"`
	ODataType                            *string `json:"@odata.type,omitempty"`
}
