package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenTypeExtensionOperationPredicate struct {
	ExtensionName *string
	Id            *string
	ODataType     *string
}

func (p OpenTypeExtensionOperationPredicate) Matches(input OpenTypeExtension) bool {

	if p.ExtensionName != nil && (input.ExtensionName == nil || *p.ExtensionName != *input.ExtensionName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
