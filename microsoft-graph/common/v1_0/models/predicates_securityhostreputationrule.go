package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostReputationRuleOperationPredicate struct {
	Description       *string
	Name              *string
	ODataType         *string
	RelatedDetailsUrl *string
}

func (p SecurityHostReputationRuleOperationPredicate) Matches(input SecurityHostReputationRule) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RelatedDetailsUrl != nil && (input.RelatedDetailsUrl == nil || *p.RelatedDetailsUrl != *input.RelatedDetailsUrl) {
		return false
	}

	return true
}
