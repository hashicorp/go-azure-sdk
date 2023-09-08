package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecentNotebookOperationPredicate struct {
	DisplayName      *string
	LastAccessedTime *string
	ODataType        *string
}

func (p RecentNotebookOperationPredicate) Matches(input RecentNotebook) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.LastAccessedTime != nil && (input.LastAccessedTime == nil || *p.LastAccessedTime != *input.LastAccessedTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
