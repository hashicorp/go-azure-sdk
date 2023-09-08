package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsServiceUserAgentOperationPredicate struct {
	ApplicationVersion *string
	HeaderValue        *string
	ODataType          *string
}

func (p CallRecordsServiceUserAgentOperationPredicate) Matches(input CallRecordsServiceUserAgent) bool {

	if p.ApplicationVersion != nil && (input.ApplicationVersion == nil || *p.ApplicationVersion != *input.ApplicationVersion) {
		return false
	}

	if p.HeaderValue != nil && (input.HeaderValue == nil || *p.HeaderValue != *input.HeaderValue) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
