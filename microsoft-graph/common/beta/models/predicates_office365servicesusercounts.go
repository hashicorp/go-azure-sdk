package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365ServicesUserCountsOperationPredicate struct {
	ExchangeActive           *int64
	ExchangeInactive         *int64
	Id                       *string
	ODataType                *string
	Office365Active          *int64
	Office365Inactive        *int64
	OneDriveActive           *int64
	OneDriveInactive         *int64
	ReportPeriod             *string
	ReportRefreshDate        *string
	SharePointActive         *int64
	SharePointInactive       *int64
	SkypeForBusinessActive   *int64
	SkypeForBusinessInactive *int64
	TeamsActive              *int64
	TeamsInactive            *int64
	YammerActive             *int64
	YammerInactive           *int64
}

func (p Office365ServicesUserCountsOperationPredicate) Matches(input Office365ServicesUserCounts) bool {

	if p.ExchangeActive != nil && (input.ExchangeActive == nil || *p.ExchangeActive != *input.ExchangeActive) {
		return false
	}

	if p.ExchangeInactive != nil && (input.ExchangeInactive == nil || *p.ExchangeInactive != *input.ExchangeInactive) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Office365Active != nil && (input.Office365Active == nil || *p.Office365Active != *input.Office365Active) {
		return false
	}

	if p.Office365Inactive != nil && (input.Office365Inactive == nil || *p.Office365Inactive != *input.Office365Inactive) {
		return false
	}

	if p.OneDriveActive != nil && (input.OneDriveActive == nil || *p.OneDriveActive != *input.OneDriveActive) {
		return false
	}

	if p.OneDriveInactive != nil && (input.OneDriveInactive == nil || *p.OneDriveInactive != *input.OneDriveInactive) {
		return false
	}

	if p.ReportPeriod != nil && (input.ReportPeriod == nil || *p.ReportPeriod != *input.ReportPeriod) {
		return false
	}

	if p.ReportRefreshDate != nil && (input.ReportRefreshDate == nil || *p.ReportRefreshDate != *input.ReportRefreshDate) {
		return false
	}

	if p.SharePointActive != nil && (input.SharePointActive == nil || *p.SharePointActive != *input.SharePointActive) {
		return false
	}

	if p.SharePointInactive != nil && (input.SharePointInactive == nil || *p.SharePointInactive != *input.SharePointInactive) {
		return false
	}

	if p.SkypeForBusinessActive != nil && (input.SkypeForBusinessActive == nil || *p.SkypeForBusinessActive != *input.SkypeForBusinessActive) {
		return false
	}

	if p.SkypeForBusinessInactive != nil && (input.SkypeForBusinessInactive == nil || *p.SkypeForBusinessInactive != *input.SkypeForBusinessInactive) {
		return false
	}

	if p.TeamsActive != nil && (input.TeamsActive == nil || *p.TeamsActive != *input.TeamsActive) {
		return false
	}

	if p.TeamsInactive != nil && (input.TeamsInactive == nil || *p.TeamsInactive != *input.TeamsInactive) {
		return false
	}

	if p.YammerActive != nil && (input.YammerActive == nil || *p.YammerActive != *input.YammerActive) {
		return false
	}

	if p.YammerInactive != nil && (input.YammerInactive == nil || *p.YammerInactive != *input.YammerInactive) {
		return false
	}

	return true
}
