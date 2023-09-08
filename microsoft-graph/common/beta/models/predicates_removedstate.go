package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemovedStateOperationPredicate struct {
	ODataType *string
	Reason    *string
}

func (p RemovedStateOperationPredicate) Matches(input RemovedState) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Reason != nil && (input.Reason == nil || *p.Reason != *input.Reason) {
		return false
	}

	return true
}
