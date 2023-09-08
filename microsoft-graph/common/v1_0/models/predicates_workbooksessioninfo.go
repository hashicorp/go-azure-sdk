package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookSessionInfoOperationPredicate struct {
	Id             *string
	ODataType      *string
	PersistChanges *bool
}

func (p WorkbookSessionInfoOperationPredicate) Matches(input WorkbookSessionInfo) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PersistChanges != nil && (input.PersistChanges == nil || *p.PersistChanges != *input.PersistChanges) {
		return false
	}

	return true
}
