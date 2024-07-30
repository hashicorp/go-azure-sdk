package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParseExpressionResponse struct {
	Error               *PublicError            `json:"error,omitempty"`
	EvaluationResult    *[]string               `json:"evaluationResult,omitempty"`
	EvaluationSucceeded *bool                   `json:"evaluationSucceeded,omitempty"`
	ODataType           *string                 `json:"@odata.type,omitempty"`
	ParsedExpression    *AttributeMappingSource `json:"parsedExpression,omitempty"`
	ParsingSucceeded    *bool                   `json:"parsingSucceeded,omitempty"`
}
