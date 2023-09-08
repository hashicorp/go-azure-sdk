package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptErrorOperationPredicate struct {
	Message   *string
	ODataType *string
}

func (p DeviceComplianceScriptErrorOperationPredicate) Matches(input DeviceComplianceScriptError) bool {

	if p.Message != nil && (input.Message == nil || *p.Message != *input.Message) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
