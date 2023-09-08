package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxSettingsOperationPredicate struct {
	ArchiveFolder *string
	DateFormat    *string
	ODataType     *string
	TimeFormat    *string
	TimeZone      *string
}

func (p MailboxSettingsOperationPredicate) Matches(input MailboxSettings) bool {

	if p.ArchiveFolder != nil && (input.ArchiveFolder == nil || *p.ArchiveFolder != *input.ArchiveFolder) {
		return false
	}

	if p.DateFormat != nil && (input.DateFormat == nil || *p.DateFormat != *input.DateFormat) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TimeFormat != nil && (input.TimeFormat == nil || *p.TimeFormat != *input.TimeFormat) {
		return false
	}

	if p.TimeZone != nil && (input.TimeZone == nil || *p.TimeZone != *input.TimeZone) {
		return false
	}

	return true
}
