package transfers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetailedTransferStatus struct {
	ErrorDetails   *TransferError         `json:"errorDetails,omitempty"`
	ProductId      *string                `json:"productId,omitempty"`
	ProductName    *string                `json:"productName,omitempty"`
	ProductType    *ProductType           `json:"productType,omitempty"`
	SkuDescription *string                `json:"skuDescription,omitempty"`
	TransferStatus *ProductTransferStatus `json:"transferStatus,omitempty"`
}
