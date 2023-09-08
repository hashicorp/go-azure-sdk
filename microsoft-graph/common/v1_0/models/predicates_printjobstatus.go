package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobStatusOperationPredicate struct {
	Description         *string
	IsAcquiredByPrinter *bool
	ODataType           *string
}

func (p PrintJobStatusOperationPredicate) Matches(input PrintJobStatus) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.IsAcquiredByPrinter != nil && (input.IsAcquiredByPrinter == nil || *p.IsAcquiredByPrinter != *input.IsAcquiredByPrinter) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
