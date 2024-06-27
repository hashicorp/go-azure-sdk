package transaction

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RefundTransactionDetails struct {
	AmountRefunded    *Amount `json:"amountRefunded,omitempty"`
	AmountRequested   *Amount `json:"amountRequested,omitempty"`
	RefundOperationId *string `json:"refundOperationId,omitempty"`
}
