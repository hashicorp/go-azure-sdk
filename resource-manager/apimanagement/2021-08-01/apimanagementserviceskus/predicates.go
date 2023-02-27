// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementserviceskus

type ResourceSkuResultOperationPredicate struct {
	ResourceType *string
}

func (p ResourceSkuResultOperationPredicate) Matches(input ResourceSkuResult) bool {

	if p.ResourceType != nil && (input.ResourceType == nil && *p.ResourceType != *input.ResourceType) {
		return false
	}

	return true
}
