package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageMentionOperationPredicate struct {
	Id          *int64
	MentionText *string
	ODataType   *string
}

func (p ChatMessageMentionOperationPredicate) Matches(input ChatMessageMention) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.MentionText != nil && (input.MentionText == nil || *p.MentionText != *input.MentionText) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
