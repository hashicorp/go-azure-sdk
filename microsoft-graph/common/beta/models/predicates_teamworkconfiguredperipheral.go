package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkConfiguredPeripheralOperationPredicate struct {
	IsOptional *bool
	ODataType  *string
}

func (p TeamworkConfiguredPeripheralOperationPredicate) Matches(input TeamworkConfiguredPeripheral) bool {

	if p.IsOptional != nil && (input.IsOptional == nil || *p.IsOptional != *input.IsOptional) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
