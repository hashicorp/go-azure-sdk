// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package availableskus

type DataBoxEdgeSkuOperationPredicate struct {
	Family       *string
	Kind         *string
	ResourceType *string
	Size         *string
}

func (p DataBoxEdgeSkuOperationPredicate) Matches(input DataBoxEdgeSku) bool {

	if p.Family != nil && (input.Family == nil && *p.Family != *input.Family) {
		return false
	}

	if p.Kind != nil && (input.Kind == nil && *p.Kind != *input.Kind) {
		return false
	}

	if p.ResourceType != nil && (input.ResourceType == nil && *p.ResourceType != *input.ResourceType) {
		return false
	}

	if p.Size != nil && (input.Size == nil && *p.Size != *input.Size) {
		return false
	}

	return true
}
