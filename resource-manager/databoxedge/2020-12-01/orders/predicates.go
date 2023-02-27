// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orders

type OrderOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p OrderOperationPredicate) Matches(input Order) bool {

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
