package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestorManagerOperationPredicate struct {
	ManagerLevel *int64
	ODataType    *string
}

func (p RequestorManagerOperationPredicate) Matches(input RequestorManager) bool {

	if p.ManagerLevel != nil && (input.ManagerLevel == nil || *p.ManagerLevel != *input.ManagerLevel) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
