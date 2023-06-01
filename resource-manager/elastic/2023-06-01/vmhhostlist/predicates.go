package vmhhostlist

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMResourcesOperationPredicate struct {
	VMResourceId *string
}

func (p VMResourcesOperationPredicate) Matches(input VMResources) bool {

	if p.VMResourceId != nil && (input.VMResourceId == nil && *p.VMResourceId != *input.VMResourceId) {
		return false
	}

	return true
}
