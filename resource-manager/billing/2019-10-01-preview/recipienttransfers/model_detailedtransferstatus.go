package recipienttransfers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetailedTransferStatus struct {
	ErrorDetails   *Error                 `json:"errorDetails,omitempty"`
	ProductId      *string                `json:"productId,omitempty"`
	ProductType    *ProductType           `json:"productType,omitempty"`
	TransferStatus *ProductTransferStatus `json:"transferStatus,omitempty"`
}
