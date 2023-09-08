package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaOperationPredicate struct {
	IsTranscriptionShown *bool
	ODataType            *string
}

func (p MediaOperationPredicate) Matches(input Media) bool {

	if p.IsTranscriptionShown != nil && (input.IsTranscriptionShown == nil || *p.IsTranscriptionShown != *input.IsTranscriptionShown) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
