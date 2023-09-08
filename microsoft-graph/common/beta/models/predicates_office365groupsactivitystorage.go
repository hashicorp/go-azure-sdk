package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365GroupsActivityStorageOperationPredicate struct {
	Id                        *string
	MailboxStorageUsedInBytes *int64
	ODataType                 *string
	ReportDate                *string
	ReportPeriod              *string
	ReportRefreshDate         *string
	SiteStorageUsedInBytes    *int64
}

func (p Office365GroupsActivityStorageOperationPredicate) Matches(input Office365GroupsActivityStorage) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.MailboxStorageUsedInBytes != nil && (input.MailboxStorageUsedInBytes == nil || *p.MailboxStorageUsedInBytes != *input.MailboxStorageUsedInBytes) {
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

	if p.SiteStorageUsedInBytes != nil && (input.SiteStorageUsedInBytes == nil || *p.SiteStorageUsedInBytes != *input.SiteStorageUsedInBytes) {
		return false
	}

	return true
}
