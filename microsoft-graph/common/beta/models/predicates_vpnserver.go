package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnServerOperationPredicate struct {
	Address         *string
	Description     *string
	IsDefaultServer *bool
	ODataType       *string
}

func (p VpnServerOperationPredicate) Matches(input VpnServer) bool {

	if p.Address != nil && (input.Address == nil || *p.Address != *input.Address) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.IsDefaultServer != nil && (input.IsDefaultServer == nil || *p.IsDefaultServer != *input.IsDefaultServer) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
