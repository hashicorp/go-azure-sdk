package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelDetailsOperationPredicate struct {
	Color       *string
	Description *string
	Id          *string
	IsActive    *bool
	Name        *string
	ODataType   *string
	Sensitivity *int64
	Tooltip     *string
}

func (p LabelDetailsOperationPredicate) Matches(input LabelDetails) bool {

	if p.Color != nil && (input.Color == nil || *p.Color != *input.Color) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsActive != nil && (input.IsActive == nil || *p.IsActive != *input.IsActive) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Sensitivity != nil && (input.Sensitivity == nil || *p.Sensitivity != *input.Sensitivity) {
		return false
	}

	if p.Tooltip != nil && (input.Tooltip == nil || *p.Tooltip != *input.Tooltip) {
		return false
	}

	return true
}
