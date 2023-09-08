package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryExportOperationOperationPredicate struct {
	CompletedDateTime *string
	CreatedDateTime   *string
	Description       *string
	Id                *string
	ODataType         *string
	OutputName        *string
	PercentProgress   *int64
}

func (p SecurityEdiscoveryExportOperationOperationPredicate) Matches(input SecurityEdiscoveryExportOperation) bool {

	if p.CompletedDateTime != nil && (input.CompletedDateTime == nil || *p.CompletedDateTime != *input.CompletedDateTime) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OutputName != nil && (input.OutputName == nil || *p.OutputName != *input.OutputName) {
		return false
	}

	if p.PercentProgress != nil && (input.PercentProgress == nil || *p.PercentProgress != *input.PercentProgress) {
		return false
	}

	return true
}
