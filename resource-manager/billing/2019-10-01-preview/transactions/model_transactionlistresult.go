package transactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransactionListResult struct {
	NextLink *string        `json:"nextLink,omitempty"`
	Value    *[]Transaction `json:"value,omitempty"`
}
