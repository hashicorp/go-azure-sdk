package products

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransferProductRequestProperties struct {
	DestinationBillingProfileId *string `json:"destinationBillingProfileId,omitempty"`
	DestinationInvoiceSectionId *string `json:"destinationInvoiceSectionId,omitempty"`
}
