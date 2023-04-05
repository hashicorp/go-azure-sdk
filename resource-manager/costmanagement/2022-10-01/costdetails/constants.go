package costdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostDetailsDataFormat string

const (
	CostDetailsDataFormatCsv CostDetailsDataFormat = "Csv"
)

func PossibleValuesForCostDetailsDataFormat() []string {
	return []string{
		string(CostDetailsDataFormatCsv),
	}
}

type CostDetailsMetricType string

const (
	CostDetailsMetricTypeActualCost    CostDetailsMetricType = "ActualCost"
	CostDetailsMetricTypeAmortizedCost CostDetailsMetricType = "AmortizedCost"
)

func PossibleValuesForCostDetailsMetricType() []string {
	return []string{
		string(CostDetailsMetricTypeActualCost),
		string(CostDetailsMetricTypeAmortizedCost),
	}
}

type CostDetailsStatusType string

const (
	CostDetailsStatusTypeCompleted   CostDetailsStatusType = "Completed"
	CostDetailsStatusTypeFailed      CostDetailsStatusType = "Failed"
	CostDetailsStatusTypeNoDataFound CostDetailsStatusType = "NoDataFound"
)

func PossibleValuesForCostDetailsStatusType() []string {
	return []string{
		string(CostDetailsStatusTypeCompleted),
		string(CostDetailsStatusTypeFailed),
		string(CostDetailsStatusTypeNoDataFound),
	}
}
