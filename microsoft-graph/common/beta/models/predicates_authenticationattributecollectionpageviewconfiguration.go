package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionPageViewConfigurationOperationPredicate struct {
	Description *string
	ODataType   *string
	Title       *string
}

func (p AuthenticationAttributeCollectionPageViewConfigurationOperationPredicate) Matches(input AuthenticationAttributeCollectionPageViewConfiguration) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	return true
}
