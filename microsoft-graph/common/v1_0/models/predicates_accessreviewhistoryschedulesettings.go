package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryScheduleSettingsOperationPredicate struct {
	ODataType   *string
	ReportRange *string
}

func (p AccessReviewHistoryScheduleSettingsOperationPredicate) Matches(input AccessReviewHistoryScheduleSettings) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReportRange != nil && (input.ReportRange == nil || *p.ReportRange != *input.ReportRange) {
		return false
	}

	return true
}
