package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationErrorOperationPredicate struct {
	EntryType            *string
	ErrorCode            *string
	ErrorMessage         *string
	Id                   *string
	JoiningValue         *string
	ODataType            *string
	RecordedDateTime     *string
	ReportableIdentifier *string
}

func (p EducationSynchronizationErrorOperationPredicate) Matches(input EducationSynchronizationError) bool {

	if p.EntryType != nil && (input.EntryType == nil || *p.EntryType != *input.EntryType) {
		return false
	}

	if p.ErrorCode != nil && (input.ErrorCode == nil || *p.ErrorCode != *input.ErrorCode) {
		return false
	}

	if p.ErrorMessage != nil && (input.ErrorMessage == nil || *p.ErrorMessage != *input.ErrorMessage) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.JoiningValue != nil && (input.JoiningValue == nil || *p.JoiningValue != *input.JoiningValue) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecordedDateTime != nil && (input.RecordedDateTime == nil || *p.RecordedDateTime != *input.RecordedDateTime) {
		return false
	}

	if p.ReportableIdentifier != nil && (input.ReportableIdentifier == nil || *p.ReportableIdentifier != *input.ReportableIdentifier) {
		return false
	}

	return true
}
