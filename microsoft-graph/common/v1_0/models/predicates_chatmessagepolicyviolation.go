package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationOperationPredicate struct {
	JustificationText *string
	ODataType         *string
}

func (p ChatMessagePolicyViolationOperationPredicate) Matches(input ChatMessagePolicyViolation) bool {

	if p.JustificationText != nil && (input.JustificationText == nil || *p.JustificationText != *input.JustificationText) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
