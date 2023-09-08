package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationPolicyTipOperationPredicate struct {
	ComplianceUrl *string
	GeneralText   *string
	ODataType     *string
}

func (p ChatMessagePolicyViolationPolicyTipOperationPredicate) Matches(input ChatMessagePolicyViolationPolicyTip) bool {

	if p.ComplianceUrl != nil && (input.ComplianceUrl == nil || *p.ComplianceUrl != *input.ComplianceUrl) {
		return false
	}

	if p.GeneralText != nil && (input.GeneralText == nil || *p.GeneralText != *input.GeneralText) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
