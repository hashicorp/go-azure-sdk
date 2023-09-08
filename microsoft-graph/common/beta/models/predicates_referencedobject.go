package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferencedObjectOperationPredicate struct {
	ODataType            *string
	ReferencedObjectName *string
	ReferencedProperty   *string
}

func (p ReferencedObjectOperationPredicate) Matches(input ReferencedObject) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReferencedObjectName != nil && (input.ReferencedObjectName == nil || *p.ReferencedObjectName != *input.ReferencedObjectName) {
		return false
	}

	if p.ReferencedProperty != nil && (input.ReferencedProperty == nil || *p.ReferencedProperty != *input.ReferencedProperty) {
		return false
	}

	return true
}
