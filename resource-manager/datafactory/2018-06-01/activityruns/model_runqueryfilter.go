package activityruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunQueryFilter struct {
	Operand  RunQueryFilterOperand  `json:"operand"`
	Operator RunQueryFilterOperator `json:"operator"`
	Values   []string               `json:"values"`
}
