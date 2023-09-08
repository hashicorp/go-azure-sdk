package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataActivityStatisticsOperationPredicate struct {
	ActivityId  *string
	DisplayName *string
	ODataType   *string
}

func (p IndustryDataIndustryDataActivityStatisticsOperationPredicate) Matches(input IndustryDataIndustryDataActivityStatistics) bool {

	if p.ActivityId != nil && (input.ActivityId == nil || *p.ActivityId != *input.ActivityId) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
