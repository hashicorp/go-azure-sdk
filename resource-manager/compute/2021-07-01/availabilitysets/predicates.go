package availabilitysets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailabilitySetOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p AvailabilitySetOperationPredicate) Matches(input AvailabilitySet) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type ResourceSkuOperationPredicate struct {
	Family       *string
	Kind         *string
	Name         *string
	ResourceType *string
	Size         *string
	Tier         *string
}

func (p ResourceSkuOperationPredicate) Matches(input ResourceSku) bool {

	if p.Family != nil && (input.Family == nil || *p.Family != *input.Family) {
		return false
	}

	if p.Kind != nil && (input.Kind == nil || *p.Kind != *input.Kind) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ResourceType != nil && (input.ResourceType == nil || *p.ResourceType != *input.ResourceType) {
		return false
	}

	if p.Size != nil && (input.Size == nil || *p.Size != *input.Size) {
		return false
	}

	if p.Tier != nil && (input.Tier == nil || *p.Tier != *input.Tier) {
		return false
	}

	return true
}
