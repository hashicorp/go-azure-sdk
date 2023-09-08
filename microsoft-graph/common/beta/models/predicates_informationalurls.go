package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationalUrlsOperationPredicate struct {
	AppSignUpUrl                 *string
	ODataType                    *string
	SingleSignOnDocumentationUrl *string
}

func (p InformationalUrlsOperationPredicate) Matches(input InformationalUrls) bool {

	if p.AppSignUpUrl != nil && (input.AppSignUpUrl == nil || *p.AppSignUpUrl != *input.AppSignUpUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SingleSignOnDocumentationUrl != nil && (input.SingleSignOnDocumentationUrl == nil || *p.SingleSignOnDocumentationUrl != *input.SingleSignOnDocumentationUrl) {
		return false
	}

	return true
}
