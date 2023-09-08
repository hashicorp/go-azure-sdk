package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PropertyToEvaluateOperationPredicate struct {
	ODataType     *string
	PropertyName  *string
	PropertyValue *string
}

func (p PropertyToEvaluateOperationPredicate) Matches(input PropertyToEvaluate) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PropertyName != nil && (input.PropertyName == nil || *p.PropertyName != *input.PropertyName) {
		return false
	}

	if p.PropertyValue != nil && (input.PropertyValue == nil || *p.PropertyValue != *input.PropertyValue) {
		return false
	}

	return true
}
