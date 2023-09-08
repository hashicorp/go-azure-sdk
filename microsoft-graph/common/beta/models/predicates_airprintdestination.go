package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AirPrintDestinationOperationPredicate struct {
	ForceTls     *bool
	IpAddress    *string
	ODataType    *string
	Port         *int64
	ResourcePath *string
}

func (p AirPrintDestinationOperationPredicate) Matches(input AirPrintDestination) bool {

	if p.ForceTls != nil && (input.ForceTls == nil || *p.ForceTls != *input.ForceTls) {
		return false
	}

	if p.IpAddress != nil && (input.IpAddress == nil || *p.IpAddress != *input.IpAddress) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Port != nil && (input.Port == nil || *p.Port != *input.Port) {
		return false
	}

	if p.ResourcePath != nil && (input.ResourcePath == nil || *p.ResourcePath != *input.ResourcePath) {
		return false
	}

	return true
}
