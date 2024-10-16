package checkpolicyrestrictions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckRestrictionEvaluationDetails struct {
	EvaluatedExpressions *[]ExpressionEvaluationDetails `json:"evaluatedExpressions,omitempty"`
	IfNotExistsDetails   *IfNotExistsEvaluationDetails  `json:"ifNotExistsDetails,omitempty"`
	Reason               *string                        `json:"reason,omitempty"`
}
