package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainStateOperationPredicate struct {
	LastActionDateTime *string
	ODataType          *string
	Operation          *string
	Status             *string
}

func (p DomainStateOperationPredicate) Matches(input DomainState) bool {

	if p.LastActionDateTime != nil && (input.LastActionDateTime == nil || *p.LastActionDateTime != *input.LastActionDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Operation != nil && (input.Operation == nil || *p.Operation != *input.Operation) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}
