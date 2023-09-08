package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementFileDataOperationPredicate struct {
	Data      *string
	ODataType *string
}

func (p AgreementFileDataOperationPredicate) Matches(input AgreementFileData) bool {

	if p.Data != nil && (input.Data == nil || *p.Data != *input.Data) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
