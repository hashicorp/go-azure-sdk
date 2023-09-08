package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RedirectUriSettingsOperationPredicate struct {
	Index     *int64
	ODataType *string
	Uri       *string
}

func (p RedirectUriSettingsOperationPredicate) Matches(input RedirectUriSettings) bool {

	if p.Index != nil && (input.Index == nil || *p.Index != *input.Index) {
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
