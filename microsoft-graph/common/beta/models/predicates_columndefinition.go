package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ColumnDefinitionOperationPredicate struct {
	ColumnGroup         *string
	Description         *string
	DisplayName         *string
	EnforceUniqueValues *bool
	Hidden              *bool
	Id                  *string
	Indexed             *bool
	IsDeletable         *bool
	IsReorderable       *bool
	IsSealed            *bool
	Name                *string
	ODataType           *string
	PropagateChanges    *bool
	ReadOnly            *bool
	Required            *bool
}

func (p ColumnDefinitionOperationPredicate) Matches(input ColumnDefinition) bool {

	if p.ColumnGroup != nil && (input.ColumnGroup == nil || *p.ColumnGroup != *input.ColumnGroup) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EnforceUniqueValues != nil && (input.EnforceUniqueValues == nil || *p.EnforceUniqueValues != *input.EnforceUniqueValues) {
		return false
	}

	if p.Hidden != nil && (input.Hidden == nil || *p.Hidden != *input.Hidden) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Indexed != nil && (input.Indexed == nil || *p.Indexed != *input.Indexed) {
		return false
	}

	if p.IsDeletable != nil && (input.IsDeletable == nil || *p.IsDeletable != *input.IsDeletable) {
		return false
	}

	if p.IsReorderable != nil && (input.IsReorderable == nil || *p.IsReorderable != *input.IsReorderable) {
		return false
	}

	if p.IsSealed != nil && (input.IsSealed == nil || *p.IsSealed != *input.IsSealed) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PropagateChanges != nil && (input.PropagateChanges == nil || *p.PropagateChanges != *input.PropagateChanges) {
		return false
	}

	if p.ReadOnly != nil && (input.ReadOnly == nil || *p.ReadOnly != *input.ReadOnly) {
		return false
	}

	if p.Required != nil && (input.Required == nil || *p.Required != *input.Required) {
		return false
	}

	return true
}
