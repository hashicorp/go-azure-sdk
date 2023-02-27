// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package trigger

type TriggerOperationPredicate struct {
}

func (p TriggerOperationPredicate) Matches(input Trigger) bool {

	return true
}
