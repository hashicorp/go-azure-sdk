package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressionEvaluationDetailsOperationPredicate struct {
	Expression       *string
	ExpressionResult *bool
	ODataType        *string
}

func (p ExpressionEvaluationDetailsOperationPredicate) Matches(input ExpressionEvaluationDetails) bool {

	if p.Expression != nil && (input.Expression == nil || *p.Expression != *input.Expression) {
		return false
	}

	if p.ExpressionResult != nil && (input.ExpressionResult == nil || *p.ExpressionResult != *input.ExpressionResult) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
