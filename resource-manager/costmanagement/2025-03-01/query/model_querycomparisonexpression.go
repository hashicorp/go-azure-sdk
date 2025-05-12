package query

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryComparisonExpression struct {
	Name     string            `json:"name"`
	Operator QueryOperatorType `json:"operator"`
	Values   []string          `json:"values"`
}
