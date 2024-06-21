package replicationjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobErrorDetails struct {
	CreationTime         *string        `json:"creationTime,omitempty"`
	ErrorLevel           *string        `json:"errorLevel,omitempty"`
	ProviderErrorDetails *ProviderError `json:"providerErrorDetails,omitempty"`
	ServiceErrorDetails  *ServiceError  `json:"serviceErrorDetails,omitempty"`
	TaskId               *string        `json:"taskId,omitempty"`
}
