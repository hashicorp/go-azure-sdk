package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskDefinitionOperationPredicate struct {
	ContinueOnError *bool
	Description     *string
	DisplayName     *string
	Id              *string
	ODataType       *string
	Version         *int64
}

func (p IdentityGovernanceTaskDefinitionOperationPredicate) Matches(input IdentityGovernanceTaskDefinition) bool {

	if p.ContinueOnError != nil && (input.ContinueOnError == nil || *p.ContinueOnError != *input.ContinueOnError) {
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

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
