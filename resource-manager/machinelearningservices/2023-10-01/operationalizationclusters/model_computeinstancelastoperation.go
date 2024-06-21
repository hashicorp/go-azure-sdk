package operationalizationclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeInstanceLastOperation struct {
	OperationName    *OperationName    `json:"operationName,omitempty"`
	OperationStatus  *OperationStatus  `json:"operationStatus,omitempty"`
	OperationTime    *string           `json:"operationTime,omitempty"`
	OperationTrigger *OperationTrigger `json:"operationTrigger,omitempty"`
}
