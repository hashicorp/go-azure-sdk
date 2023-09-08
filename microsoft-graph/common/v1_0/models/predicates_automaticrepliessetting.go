package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticRepliesSettingOperationPredicate struct {
	ExternalReplyMessage *string
	InternalReplyMessage *string
	ODataType            *string
}

func (p AutomaticRepliesSettingOperationPredicate) Matches(input AutomaticRepliesSetting) bool {

	if p.ExternalReplyMessage != nil && (input.ExternalReplyMessage == nil || *p.ExternalReplyMessage != *input.ExternalReplyMessage) {
		return false
	}

	if p.InternalReplyMessage != nil && (input.InternalReplyMessage == nil || *p.InternalReplyMessage != *input.InternalReplyMessage) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
