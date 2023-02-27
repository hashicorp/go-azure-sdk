// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertrules

type AlertRuleOperationPredicate struct {
}

func (p AlertRuleOperationPredicate) Matches(input AlertRule) bool {

	return true
}
