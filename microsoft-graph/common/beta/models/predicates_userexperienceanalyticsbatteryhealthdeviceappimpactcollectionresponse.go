package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthDeviceAppImpactCollectionResponseOperationPredicate struct {
	ODataCount    *int64
	ODataNextLink *string
}

func (p UserExperienceAnalyticsBatteryHealthDeviceAppImpactCollectionResponseOperationPredicate) Matches(input UserExperienceAnalyticsBatteryHealthDeviceAppImpactCollectionResponse) bool {

	if p.ODataCount != nil && (input.ODataCount == nil || *p.ODataCount != *input.ODataCount) {
		return false
	}

	if p.ODataNextLink != nil && (input.ODataNextLink == nil || *p.ODataNextLink != *input.ODataNextLink) {
		return false
	}

	return true
}
