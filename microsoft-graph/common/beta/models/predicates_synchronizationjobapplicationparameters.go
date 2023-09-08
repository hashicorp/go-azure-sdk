package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJobApplicationParametersOperationPredicate struct {
	ODataType *string
	RuleId    *string
}

func (p SynchronizationJobApplicationParametersOperationPredicate) Matches(input SynchronizationJobApplicationParameters) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RuleId != nil && (input.RuleId == nil || *p.RuleId != *input.RuleId) {
		return false
	}

	return true
}
