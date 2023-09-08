package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppHostedMediaConfigOperationPredicate struct {
	Blob      *string
	ODataType *string
}

func (p AppHostedMediaConfigOperationPredicate) Matches(input AppHostedMediaConfig) bool {

	if p.Blob != nil && (input.Blob == nil || *p.Blob != *input.Blob) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
