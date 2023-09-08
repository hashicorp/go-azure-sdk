package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LastSignInOperationPredicate struct {
	InteractiveDateTime    *string
	NonInteractiveDateTime *string
	ODataType              *string
}

func (p LastSignInOperationPredicate) Matches(input LastSignIn) bool {

	if p.InteractiveDateTime != nil && (input.InteractiveDateTime == nil || *p.InteractiveDateTime != *input.InteractiveDateTime) {
		return false
	}

	if p.NonInteractiveDateTime != nil && (input.NonInteractiveDateTime == nil || *p.NonInteractiveDateTime != *input.NonInteractiveDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
