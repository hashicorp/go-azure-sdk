package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationBandwidthAbsolute struct {
	MaximumDownloadBandwidthInKilobytesPerSecond *int64  `json:"maximumDownloadBandwidthInKilobytesPerSecond,omitempty"`
	MaximumUploadBandwidthInKilobytesPerSecond   *int64  `json:"maximumUploadBandwidthInKilobytesPerSecond,omitempty"`
	ODataType                                    *string `json:"@odata.type,omitempty"`
}
