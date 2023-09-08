package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ToneInfoOperationPredicate struct {
	ODataType  *string
	SequenceId *int64
}

func (p ToneInfoOperationPredicate) Matches(input ToneInfo) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SequenceId != nil && (input.SequenceId == nil || *p.SequenceId != *input.SequenceId) {
		return false
	}

	return true
}
