package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchemaExtensionOperationPredicate struct {
	Description *string
	Id          *string
	ODataType   *string
	Owner       *string
	Status      *string
}

func (p SchemaExtensionOperationPredicate) Matches(input SchemaExtension) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Owner != nil && (input.Owner == nil || *p.Owner != *input.Owner) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}
