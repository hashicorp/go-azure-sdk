package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitedUserMessageInfoOperationPredicate struct {
	CustomizedMessageBody *string
	MessageLanguage       *string
	ODataType             *string
}

func (p InvitedUserMessageInfoOperationPredicate) Matches(input InvitedUserMessageInfo) bool {

	if p.CustomizedMessageBody != nil && (input.CustomizedMessageBody == nil || *p.CustomizedMessageBody != *input.CustomizedMessageBody) {
		return false
	}

	if p.MessageLanguage != nil && (input.MessageLanguage == nil || *p.MessageLanguage != *input.MessageLanguage) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
