package policystates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressionEvaluationDetails struct {
	Expression      *string      `json:"expression,omitempty"`
	ExpressionKind  *string      `json:"expressionKind,omitempty"`
	ExpressionValue *interface{} `json:"expressionValue,omitempty"`
	Operator        *string      `json:"operator,omitempty"`
	Path            *string      `json:"path,omitempty"`
	Result          *string      `json:"result,omitempty"`
	TargetValue     *interface{} `json:"targetValue,omitempty"`
}
