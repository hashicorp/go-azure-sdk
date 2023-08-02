package skus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabServicesSkuOperationPredicate struct {
	Family       *string
	Name         *string
	ResourceType *string
	Size         *string
}

func (p LabServicesSkuOperationPredicate) Matches(input LabServicesSku) bool {

	if p.Family != nil && (input.Family == nil || *p.Family != *input.Family) {
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

	return true
}
