package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenerateDetailedCostReportDefinition struct {
	BillingPeriod *string                               `json:"billingPeriod,omitempty"`
	CustomerId    *string                               `json:"customerId,omitempty"`
	InvoiceId     *string                               `json:"invoiceId,omitempty"`
	Metric        *GenerateDetailedCostReportMetricType `json:"metric,omitempty"`
	TimePeriod    *GenerateDetailedCostReportTimePeriod `json:"timePeriod,omitempty"`
}
