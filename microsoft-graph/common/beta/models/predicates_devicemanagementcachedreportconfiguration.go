package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCachedReportConfigurationOperationPredicate struct {
	ExpirationDateTime  *string
	Filter              *string
	Id                  *string
	LastRefreshDateTime *string
	Metadata            *string
	ODataType           *string
	ReportName          *string
}

func (p DeviceManagementCachedReportConfigurationOperationPredicate) Matches(input DeviceManagementCachedReportConfiguration) bool {

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.Filter != nil && (input.Filter == nil || *p.Filter != *input.Filter) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastRefreshDateTime != nil && (input.LastRefreshDateTime == nil || *p.LastRefreshDateTime != *input.LastRefreshDateTime) {
		return false
	}

	if p.Metadata != nil && (input.Metadata == nil || *p.Metadata != *input.Metadata) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReportName != nil && (input.ReportName == nil || *p.ReportName != *input.ReportName) {
		return false
	}

	return true
}
