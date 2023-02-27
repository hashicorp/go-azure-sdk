// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hybridrunbookworkergroup

type HybridRunbookWorkerGroupOperationPredicate struct {
	Id   *string
	Name *string
}

func (p HybridRunbookWorkerGroupOperationPredicate) Matches(input HybridRunbookWorkerGroup) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	return true
}
