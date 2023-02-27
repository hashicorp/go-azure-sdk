// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobstream

type JobStreamOperationPredicate struct {
	Id *string
}

func (p JobStreamOperationPredicate) Matches(input JobStream) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	return true
}
