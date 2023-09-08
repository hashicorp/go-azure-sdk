package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterShareViewpointOperationPredicate struct {
	LastUsedDateTime *string
	ODataType        *string
}

func (p PrinterShareViewpointOperationPredicate) Matches(input PrinterShareViewpoint) bool {

	if p.LastUsedDateTime != nil && (input.LastUsedDateTime == nil || *p.LastUsedDateTime != *input.LastUsedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
