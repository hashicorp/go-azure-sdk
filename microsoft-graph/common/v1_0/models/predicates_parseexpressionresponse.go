package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParseExpressionResponseOperationPredicate struct {
	EvaluationSucceeded *bool
	ODataType           *string
	ParsingSucceeded    *bool
}

func (p ParseExpressionResponseOperationPredicate) Matches(input ParseExpressionResponse) bool {

	if p.EvaluationSucceeded != nil && (input.EvaluationSucceeded == nil || *p.EvaluationSucceeded != *input.EvaluationSucceeded) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ParsingSucceeded != nil && (input.ParsingSucceeded == nil || *p.ParsingSucceeded != *input.ParsingSucceeded) {
		return false
	}

	return true
}
