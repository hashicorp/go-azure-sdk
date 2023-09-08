package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceRuleSettingOperationPredicate struct {
	ODataType      *string
	RuleIdentifier *string
	Setting        *string
}

func (p GovernanceRuleSettingOperationPredicate) Matches(input GovernanceRuleSetting) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RuleIdentifier != nil && (input.RuleIdentifier == nil || *p.RuleIdentifier != *input.RuleIdentifier) {
		return false
	}

	if p.Setting != nil && (input.Setting == nil || *p.Setting != *input.Setting) {
		return false
	}

	return true
}
