package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10AssociatedAppsOperationPredicate struct {
	Identifier *string
	ODataType  *string
}

func (p Windows10AssociatedAppsOperationPredicate) Matches(input Windows10AssociatedApps) bool {

	if p.Identifier != nil && (input.Identifier == nil || *p.Identifier != *input.Identifier) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
