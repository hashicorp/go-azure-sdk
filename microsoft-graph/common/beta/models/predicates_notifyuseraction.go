package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotifyUserActionOperationPredicate struct {
	ActionLastModifiedDateTime *string
	EmailText                  *string
	ODataType                  *string
	PolicyTip                  *string
}

func (p NotifyUserActionOperationPredicate) Matches(input NotifyUserAction) bool {

	if p.ActionLastModifiedDateTime != nil && (input.ActionLastModifiedDateTime == nil || *p.ActionLastModifiedDateTime != *input.ActionLastModifiedDateTime) {
		return false
	}

	if p.EmailText != nil && (input.EmailText == nil || *p.EmailText != *input.EmailText) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PolicyTip != nil && (input.PolicyTip == nil || *p.PolicyTip != *input.PolicyTip) {
		return false
	}

	return true
}
