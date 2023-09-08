package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactFolderOperationPredicate struct {
	DisplayName    *string
	Id             *string
	ODataType      *string
	ParentFolderId *string
}

func (p ContactFolderOperationPredicate) Matches(input ContactFolder) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ParentFolderId != nil && (input.ParentFolderId == nil || *p.ParentFolderId != *input.ParentFolderId) {
		return false
	}

	return true
}
