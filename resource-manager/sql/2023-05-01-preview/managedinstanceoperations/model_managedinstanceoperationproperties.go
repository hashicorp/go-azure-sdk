package managedinstanceoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceOperationProperties struct {
	Description             *string                                 `json:"description,omitempty"`
	ErrorCode               *int64                                  `json:"errorCode,omitempty"`
	ErrorDescription        *string                                 `json:"errorDescription,omitempty"`
	ErrorSeverity           *int64                                  `json:"errorSeverity,omitempty"`
	EstimatedCompletionTime *string                                 `json:"estimatedCompletionTime,omitempty"`
	IsCancellable           *bool                                   `json:"isCancellable,omitempty"`
	IsUserError             *bool                                   `json:"isUserError,omitempty"`
	ManagedInstanceName     *string                                 `json:"managedInstanceName,omitempty"`
	Operation               *string                                 `json:"operation,omitempty"`
	OperationFriendlyName   *string                                 `json:"operationFriendlyName,omitempty"`
	OperationParameters     *ManagedInstanceOperationParametersPair `json:"operationParameters,omitempty"`
	OperationSteps          *ManagedInstanceOperationSteps          `json:"operationSteps,omitempty"`
	PercentComplete         *int64                                  `json:"percentComplete,omitempty"`
	StartTime               *string                                 `json:"startTime,omitempty"`
	State                   *ManagementOperationState               `json:"state,omitempty"`
}
