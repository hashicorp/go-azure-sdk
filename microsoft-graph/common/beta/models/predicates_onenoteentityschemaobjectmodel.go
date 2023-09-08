package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteEntitySchemaObjectModelOperationPredicate struct {
	CreatedDateTime *string
	Id              *string
	ODataType       *string
	Self            *string
}

func (p OnenoteEntitySchemaObjectModelOperationPredicate) Matches(input OnenoteEntitySchemaObjectModel) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Self != nil && (input.Self == nil || *p.Self != *input.Self) {
		return false
	}

	return true
}
