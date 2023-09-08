package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineContributingPolicyOperationPredicate struct {
	DisplayName *string
	ODataType   *string
	SourceId    *string
}

func (p SecurityBaselineContributingPolicyOperationPredicate) Matches(input SecurityBaselineContributingPolicy) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SourceId != nil && (input.SourceId == nil || *p.SourceId != *input.SourceId) {
		return false
	}

	return true
}
