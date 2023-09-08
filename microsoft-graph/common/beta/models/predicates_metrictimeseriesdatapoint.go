package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricTimeSeriesDataPointOperationPredicate struct {
	DateTime  *string
	ODataType *string
	Value     *int64
}

func (p MetricTimeSeriesDataPointOperationPredicate) Matches(input MetricTimeSeriesDataPoint) bool {

	if p.DateTime != nil && (input.DateTime == nil || *p.DateTime != *input.DateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
