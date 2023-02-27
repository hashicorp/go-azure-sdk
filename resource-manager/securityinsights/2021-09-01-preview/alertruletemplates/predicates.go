// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertruletemplates

type AlertRuleTemplateOperationPredicate struct {
}

func (p AlertRuleTemplateOperationPredicate) Matches(input AlertRuleTemplate) bool {

	return true
}
