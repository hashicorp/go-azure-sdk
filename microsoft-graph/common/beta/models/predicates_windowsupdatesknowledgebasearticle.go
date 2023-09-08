package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesKnowledgeBaseArticleOperationPredicate struct {
	Id        *string
	ODataType *string
	Url       *string
}

func (p WindowsUpdatesKnowledgeBaseArticleOperationPredicate) Matches(input WindowsUpdatesKnowledgeBaseArticle) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	return true
}
