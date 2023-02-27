// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package skus

type ResourceSkuOperationPredicate struct {
	Name         *string
	ResourceType *string
}

func (p ResourceSkuOperationPredicate) Matches(input ResourceSku) bool {

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.ResourceType != nil && (input.ResourceType == nil && *p.ResourceType != *input.ResourceType) {
		return false
	}

	return true
}
