package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCBulkRestoreOperationPredicate struct {
	CreatedDateTime      *string
	DisplayName          *string
	Id                   *string
	ODataType            *string
	RestorePointDateTime *string
}

func (p CloudPCBulkRestoreOperationPredicate) Matches(input CloudPCBulkRestore) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RestorePointDateTime != nil && (input.RestorePointDateTime == nil || *p.RestorePointDateTime != *input.RestorePointDateTime) {
		return false
	}

	return true
}
