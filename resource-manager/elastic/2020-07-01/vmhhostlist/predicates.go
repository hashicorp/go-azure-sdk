// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmhhostlist

type VMResourcesOperationPredicate struct {
	VMResourceId *string
}

func (p VMResourcesOperationPredicate) Matches(input VMResources) bool {

	if p.VMResourceId != nil && (input.VMResourceId == nil && *p.VMResourceId != *input.VMResourceId) {
		return false
	}

	return true
}
