// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package contenttypecontentitem

type ContentItemContractOperationPredicate struct {
	Id         *string
	Name       *string
	Properties *interface{}
	Type       *string
}

func (p ContentItemContractOperationPredicate) Matches(input ContentItemContract) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Properties != nil && (input.Properties == nil && *p.Properties != *input.Properties) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}
