package query

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryFilter struct {
	And        *[]QueryFilter             `json:"and,omitempty"`
	Dimensions *QueryComparisonExpression `json:"dimensions,omitempty"`
	Or         *[]QueryFilter             `json:"or,omitempty"`
	Tags       *QueryComparisonExpression `json:"tags,omitempty"`
}
