package costdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenerateCostDetailsReportRequestDefinition struct {
	BillingPeriod *string                `json:"billingPeriod,omitempty"`
	InvoiceId     *string                `json:"invoiceId,omitempty"`
	Metric        *CostDetailsMetricType `json:"metric,omitempty"`
	TimePeriod    *CostDetailsTimePeriod `json:"timePeriod"`
}
