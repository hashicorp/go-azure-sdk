package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWindows10DevicesSummary struct {
	ODataType                       *string `json:"@odata.type,omitempty"`
	UnsupportedOSversionDeviceCount *int64  `json:"unsupportedOSversionDeviceCount,omitempty"`
}
