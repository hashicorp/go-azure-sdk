package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCResizeValidationResultOperationPredicate struct {
	CloudPCId *string
	ODataType *string
}

func (p CloudPCResizeValidationResultOperationPredicate) Matches(input CloudPCResizeValidationResult) bool {

	if p.CloudPCId != nil && (input.CloudPCId == nil || *p.CloudPCId != *input.CloudPCId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
