package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppMetadataEntryOperationPredicate struct {
	Key       *string
	ODataType *string
	Value     *string
}

func (p AppMetadataEntryOperationPredicate) Matches(input AppMetadataEntry) bool {

	if p.Key != nil && (input.Key == nil || *p.Key != *input.Key) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
