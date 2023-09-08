package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsPropertyRuleOperationPredicate struct {
	ODataType *string
	Property  *string
}

func (p ExternalConnectorsPropertyRuleOperationPredicate) Matches(input ExternalConnectorsPropertyRule) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Property != nil && (input.Property == nil || *p.Property != *input.Property) {
		return false
	}

	return true
}
