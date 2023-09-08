package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityActionStateOperationPredicate struct {
	AppId           *string
	ODataType       *string
	UpdatedDateTime *string
	User            *string
}

func (p SecurityActionStateOperationPredicate) Matches(input SecurityActionState) bool {

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UpdatedDateTime != nil && (input.UpdatedDateTime == nil || *p.UpdatedDateTime != *input.UpdatedDateTime) {
		return false
	}

	if p.User != nil && (input.User == nil || *p.User != *input.User) {
		return false
	}

	return true
}
