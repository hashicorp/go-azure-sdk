package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveContentLocationOperationPredicate struct {
	Confidence *int64
	IdMatch    *string
	Length     *int64
	ODataType  *string
	Offset     *int64
}

func (p SensitiveContentLocationOperationPredicate) Matches(input SensitiveContentLocation) bool {

	if p.Confidence != nil && (input.Confidence == nil || *p.Confidence != *input.Confidence) {
		return false
	}

	if p.IdMatch != nil && (input.IdMatch == nil || *p.IdMatch != *input.IdMatch) {
		return false
	}

	if p.Length != nil && (input.Length == nil || *p.Length != *input.Length) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Offset != nil && (input.Offset == nil || *p.Offset != *input.Offset) {
		return false
	}

	return true
}
