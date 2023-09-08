package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365GroupsActivityFileCountsOperationPredicate struct {
	Active            *int64
	Id                *string
	ODataType         *string
	ReportDate        *string
	ReportPeriod      *string
	ReportRefreshDate *string
	Total             *int64
}

func (p Office365GroupsActivityFileCountsOperationPredicate) Matches(input Office365GroupsActivityFileCounts) bool {

	if p.Active != nil && (input.Active == nil || *p.Active != *input.Active) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReportDate != nil && (input.ReportDate == nil || *p.ReportDate != *input.ReportDate) {
		return false
	}

	if p.ReportPeriod != nil && (input.ReportPeriod == nil || *p.ReportPeriod != *input.ReportPeriod) {
		return false
	}

	if p.ReportRefreshDate != nil && (input.ReportRefreshDate == nil || *p.ReportRefreshDate != *input.ReportRefreshDate) {
		return false
	}

	if p.Total != nil && (input.Total == nil || *p.Total != *input.Total) {
		return false
	}

	return true
}
