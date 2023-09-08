package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageLocalizedContentOperationPredicate struct {
	DefaultText *string
	ODataType   *string
}

func (p AccessPackageLocalizedContentOperationPredicate) Matches(input AccessPackageLocalizedContent) bool {

	if p.DefaultText != nil && (input.DefaultText == nil || *p.DefaultText != *input.DefaultText) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
