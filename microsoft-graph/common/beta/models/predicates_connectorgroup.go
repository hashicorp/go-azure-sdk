package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorGroupOperationPredicate struct {
	Id        *string
	IsDefault *bool
	Name      *string
	ODataType *string
}

func (p ConnectorGroupOperationPredicate) Matches(input ConnectorGroup) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDefault != nil && (input.IsDefault == nil || *p.IsDefault != *input.IsDefault) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
