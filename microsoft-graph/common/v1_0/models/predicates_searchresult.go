package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchResultOperationPredicate struct {
	ODataType           *string
	OnClickTelemetryUrl *string
}

func (p SearchResultOperationPredicate) Matches(input SearchResult) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OnClickTelemetryUrl != nil && (input.OnClickTelemetryUrl == nil || *p.OnClickTelemetryUrl != *input.OnClickTelemetryUrl) {
		return false
	}

	return true
}
