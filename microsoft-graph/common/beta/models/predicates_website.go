package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebsiteOperationPredicate struct {
	Address     *string
	DisplayName *string
	ODataType   *string
}

func (p WebsiteOperationPredicate) Matches(input Website) bool {

	if p.Address != nil && (input.Address == nil || *p.Address != *input.Address) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
