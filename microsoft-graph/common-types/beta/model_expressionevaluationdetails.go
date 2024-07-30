package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressionEvaluationDetails struct {
	Expression                  *string                        `json:"expression,omitempty"`
	ExpressionEvaluationDetails *[]ExpressionEvaluationDetails `json:"expressionEvaluationDetails,omitempty"`
	ExpressionResult            *bool                          `json:"expressionResult,omitempty"`
	ODataType                   *string                        `json:"@odata.type,omitempty"`
	PropertyToEvaluate          *PropertyToEvaluate            `json:"propertyToEvaluate,omitempty"`
}
