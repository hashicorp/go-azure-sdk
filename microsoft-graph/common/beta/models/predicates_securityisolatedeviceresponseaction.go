package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIsolateDeviceResponseActionOperationPredicate struct {
	ODataType *string
}

func (p SecurityIsolateDeviceResponseActionOperationPredicate) Matches(input SecurityIsolateDeviceResponseAction) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
