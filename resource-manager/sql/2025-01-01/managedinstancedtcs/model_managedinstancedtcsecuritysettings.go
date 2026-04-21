package managedinstancedtcs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceDtcSecuritySettings struct {
	SnaLu6point2TransactionsEnabled         *bool                                                      `json:"snaLu6point2TransactionsEnabled,omitempty"`
	TransactionManagerCommunicationSettings *ManagedInstanceDtcTransactionManagerCommunicationSettings `json:"transactionManagerCommunicationSettings,omitempty"`
	XaTransactionsDefaultTimeout            *int64                                                     `json:"xaTransactionsDefaultTimeout,omitempty"`
	XaTransactionsEnabled                   *bool                                                      `json:"xaTransactionsEnabled,omitempty"`
	XaTransactionsMaximumTimeout            *int64                                                     `json:"xaTransactionsMaximumTimeout,omitempty"`
}
