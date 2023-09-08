package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InviteNewBotResponseOperationPredicate struct {
	InviteUri *string
	ODataType *string
}

func (p InviteNewBotResponseOperationPredicate) Matches(input InviteNewBotResponse) bool {

	if p.InviteUri != nil && (input.InviteUri == nil || *p.InviteUri != *input.InviteUri) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
