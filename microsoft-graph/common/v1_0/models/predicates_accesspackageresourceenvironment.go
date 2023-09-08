package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceEnvironmentOperationPredicate struct {
	CreatedDateTime      *string
	Description          *string
	DisplayName          *string
	Id                   *string
	IsDefaultEnvironment *bool
	ModifiedDateTime     *string
	ODataType            *string
	OriginId             *string
	OriginSystem         *string
}

func (p AccessPackageResourceEnvironmentOperationPredicate) Matches(input AccessPackageResourceEnvironment) bool {

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

	if p.IsDefaultEnvironment != nil && (input.IsDefaultEnvironment == nil || *p.IsDefaultEnvironment != *input.IsDefaultEnvironment) {
		return false
	}

	if p.ModifiedDateTime != nil && (input.ModifiedDateTime == nil || *p.ModifiedDateTime != *input.ModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OriginId != nil && (input.OriginId == nil || *p.OriginId != *input.OriginId) {
		return false
	}

	if p.OriginSystem != nil && (input.OriginSystem == nil || *p.OriginSystem != *input.OriginSystem) {
		return false
	}

	return true
}
