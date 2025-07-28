package skus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSkuOperationPredicate struct {
	Kind         *string
	Name         *string
	ResourceType *string
	Tier         *string
}

func (p ResourceSkuOperationPredicate) Matches(input ResourceSku) bool {

	if p.Kind != nil && (input.Kind == nil || *p.Kind != *input.Kind) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ResourceType != nil && (input.ResourceType == nil || *p.ResourceType != *input.ResourceType) {
		return false
	}

	if p.Tier != nil && (input.Tier == nil || *p.Tier != *input.Tier) {
		return false
	}

	return true
}
