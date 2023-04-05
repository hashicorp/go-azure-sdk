package usagedetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenerateDetailedCostReportMetricType string

const (
	GenerateDetailedCostReportMetricTypeActualCost    GenerateDetailedCostReportMetricType = "ActualCost"
	GenerateDetailedCostReportMetricTypeAmortizedCost GenerateDetailedCostReportMetricType = "AmortizedCost"
)

func PossibleValuesForGenerateDetailedCostReportMetricType() []string {
	return []string{
		string(GenerateDetailedCostReportMetricTypeActualCost),
		string(GenerateDetailedCostReportMetricTypeAmortizedCost),
	}
}
