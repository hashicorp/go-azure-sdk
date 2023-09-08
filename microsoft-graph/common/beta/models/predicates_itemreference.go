package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemReferenceOperationPredicate struct {
	DriveId   *string
	DriveType *string
	Id        *string
	Name      *string
	ODataType *string
	Path      *string
	ShareId   *string
	SiteId    *string
}

func (p ItemReferenceOperationPredicate) Matches(input ItemReference) bool {

	if p.DriveId != nil && (input.DriveId == nil || *p.DriveId != *input.DriveId) {
		return false
	}

	if p.DriveType != nil && (input.DriveType == nil || *p.DriveType != *input.DriveType) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Path != nil && (input.Path == nil || *p.Path != *input.Path) {
		return false
	}

	if p.ShareId != nil && (input.ShareId == nil || *p.ShareId != *input.ShareId) {
		return false
	}

	if p.SiteId != nil && (input.SiteId == nil || *p.SiteId != *input.SiteId) {
		return false
	}

	return true
}
