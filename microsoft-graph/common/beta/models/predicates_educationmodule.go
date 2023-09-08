package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationModuleOperationPredicate struct {
	CreatedDateTime      *string
	Description          *string
	DisplayName          *string
	Id                   *string
	IsPinned             *bool
	LastModifiedDateTime *string
	ODataType            *string
	ResourcesFolderUrl   *string
}

func (p EducationModuleOperationPredicate) Matches(input EducationModule) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsPinned != nil && (input.IsPinned == nil || *p.IsPinned != *input.IsPinned) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ResourcesFolderUrl != nil && (input.ResourcesFolderUrl == nil || *p.ResourcesFolderUrl != *input.ResourcesFolderUrl) {
		return false
	}

	return true
}
