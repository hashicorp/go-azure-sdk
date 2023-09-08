package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetaDataKeyValuePairOperationPredicate struct {
	Key       *string
	ODataType *string
}

func (p MetaDataKeyValuePairOperationPredicate) Matches(input MetaDataKeyValuePair) bool {

	if p.Key != nil && (input.Key == nil || *p.Key != *input.Key) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
