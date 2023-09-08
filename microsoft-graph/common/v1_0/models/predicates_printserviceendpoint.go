package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintServiceEndpointOperationPredicate struct {
	DisplayName *string
	Id          *string
	ODataType   *string
	Uri         *string
}

func (p PrintServiceEndpointOperationPredicate) Matches(input PrintServiceEndpoint) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Uri != nil && (input.Uri == nil || *p.Uri != *input.Uri) {
		return false
	}

	return true
}
