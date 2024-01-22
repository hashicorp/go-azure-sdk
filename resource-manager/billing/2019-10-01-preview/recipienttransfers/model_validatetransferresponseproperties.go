package recipienttransfers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateTransferResponseProperties struct {
	ProductId *string                       `json:"productId,omitempty"`
	Results   *[]ValidationResultProperties `json:"results,omitempty"`
	Status    *string                       `json:"status,omitempty"`
}
