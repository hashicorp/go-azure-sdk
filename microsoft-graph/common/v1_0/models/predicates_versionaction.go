package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VersionActionOperationPredicate struct {
	NewVersion *string
	ODataType  *string
}

func (p VersionActionOperationPredicate) Matches(input VersionAction) bool {

	if p.NewVersion != nil && (input.NewVersion == nil || *p.NewVersion != *input.NewVersion) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
