package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LookupColumnOperationPredicate struct {
	AllowMultipleValues   *bool
	AllowUnlimitedLength  *bool
	ColumnName            *string
	ListId                *string
	ODataType             *string
	PrimaryLookupColumnId *string
}

func (p LookupColumnOperationPredicate) Matches(input LookupColumn) bool {

	if p.AllowMultipleValues != nil && (input.AllowMultipleValues == nil || *p.AllowMultipleValues != *input.AllowMultipleValues) {
		return false
	}

	if p.AllowUnlimitedLength != nil && (input.AllowUnlimitedLength == nil || *p.AllowUnlimitedLength != *input.AllowUnlimitedLength) {
		return false
	}

	if p.ColumnName != nil && (input.ColumnName == nil || *p.ColumnName != *input.ColumnName) {
		return false
	}

	if p.ListId != nil && (input.ListId == nil || *p.ListId != *input.ListId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrimaryLookupColumnId != nil && (input.PrimaryLookupColumnId == nil || *p.PrimaryLookupColumnId != *input.PrimaryLookupColumnId) {
		return false
	}

	return true
}
