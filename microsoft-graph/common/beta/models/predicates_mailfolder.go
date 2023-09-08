package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderOperationPredicate struct {
	ChildFolderCount *int64
	DisplayName      *string
	Id               *string
	IsHidden         *bool
	ODataType        *string
	ParentFolderId   *string
	TotalItemCount   *int64
	UnreadItemCount  *int64
	WellKnownName    *string
}

func (p MailFolderOperationPredicate) Matches(input MailFolder) bool {

	if p.ChildFolderCount != nil && (input.ChildFolderCount == nil || *p.ChildFolderCount != *input.ChildFolderCount) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsHidden != nil && (input.IsHidden == nil || *p.IsHidden != *input.IsHidden) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ParentFolderId != nil && (input.ParentFolderId == nil || *p.ParentFolderId != *input.ParentFolderId) {
		return false
	}

	if p.TotalItemCount != nil && (input.TotalItemCount == nil || *p.TotalItemCount != *input.TotalItemCount) {
		return false
	}

	if p.UnreadItemCount != nil && (input.UnreadItemCount == nil || *p.UnreadItemCount != *input.UnreadItemCount) {
		return false
	}

	if p.WellKnownName != nil && (input.WellKnownName == nil || *p.WellKnownName != *input.WellKnownName) {
		return false
	}

	return true
}
