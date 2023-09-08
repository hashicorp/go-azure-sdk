package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionPageOperationPredicate struct {
	CustomStringsFileId *string
	ODataType           *string
}

func (p AuthenticationAttributeCollectionPageOperationPredicate) Matches(input AuthenticationAttributeCollectionPage) bool {

	if p.CustomStringsFileId != nil && (input.CustomStringsFileId == nil || *p.CustomStringsFileId != *input.CustomStringsFileId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
