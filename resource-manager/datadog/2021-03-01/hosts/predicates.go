// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hosts

type DatadogHostOperationPredicate struct {
	Name *string
}

func (p DatadogHostOperationPredicate) Matches(input DatadogHost) bool {

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	return true
}
