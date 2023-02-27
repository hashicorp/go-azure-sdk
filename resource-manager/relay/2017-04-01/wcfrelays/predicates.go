// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wcfrelays

type AuthorizationRuleOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p AuthorizationRuleOperationPredicate) Matches(input AuthorizationRule) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}

type WcfRelayOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p WcfRelayOperationPredicate) Matches(input WcfRelay) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}
