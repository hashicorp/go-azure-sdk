package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceResourceOperationPredicate struct {
	DisplayName        *string
	ExternalId         *string
	Id                 *string
	ODataType          *string
	RegisteredDateTime *string
	RegisteredRoot     *string
	Status             *string
	Type               *string
}

func (p GovernanceResourceOperationPredicate) Matches(input GovernanceResource) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ExternalId != nil && (input.ExternalId == nil || *p.ExternalId != *input.ExternalId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RegisteredDateTime != nil && (input.RegisteredDateTime == nil || *p.RegisteredDateTime != *input.RegisteredDateTime) {
		return false
	}

	if p.RegisteredRoot != nil && (input.RegisteredRoot == nil || *p.RegisteredRoot != *input.RegisteredRoot) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
