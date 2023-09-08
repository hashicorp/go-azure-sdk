package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProtectAdhocActionOperationPredicate struct {
	ODataType *string
}

func (p SecurityProtectAdhocActionOperationPredicate) Matches(input SecurityProtectAdhocAction) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
