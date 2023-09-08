package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWindows10DevicesSummaryOperationPredicate struct {
	ODataType                       *string
	UnsupportedOSversionDeviceCount *int64
}

func (p UserExperienceAnalyticsWindows10DevicesSummaryOperationPredicate) Matches(input UserExperienceAnalyticsWindows10DevicesSummary) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UnsupportedOSversionDeviceCount != nil && (input.UnsupportedOSversionDeviceCount == nil || *p.UnsupportedOSversionDeviceCount != *input.UnsupportedOSversionDeviceCount) {
		return false
	}

	return true
}
