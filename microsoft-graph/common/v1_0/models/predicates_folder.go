package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FolderOperationPredicate struct {
	ChildCount *int64
	ODataType  *string
}

func (p FolderOperationPredicate) Matches(input Folder) bool {

	if p.ChildCount != nil && (input.ChildCount == nil || *p.ChildCount != *input.ChildCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
