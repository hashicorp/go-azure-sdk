package invoices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvoiceListResult struct {
	NextLink   *string    `json:"nextLink,omitempty"`
	TotalCount *float64   `json:"totalCount,omitempty"`
	Value      *[]Invoice `json:"value,omitempty"`
}
