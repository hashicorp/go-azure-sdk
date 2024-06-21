package replicationevents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthError struct {
	CreationTimeUtc              *string                           `json:"creationTimeUtc,omitempty"`
	CustomerResolvability        *HealthErrorCustomerResolvability `json:"customerResolvability,omitempty"`
	EntityId                     *string                           `json:"entityId,omitempty"`
	ErrorCategory                *string                           `json:"errorCategory,omitempty"`
	ErrorCode                    *string                           `json:"errorCode,omitempty"`
	ErrorId                      *string                           `json:"errorId,omitempty"`
	ErrorLevel                   *string                           `json:"errorLevel,omitempty"`
	ErrorMessage                 *string                           `json:"errorMessage,omitempty"`
	ErrorSource                  *string                           `json:"errorSource,omitempty"`
	ErrorType                    *string                           `json:"errorType,omitempty"`
	InnerHealthErrors            *[]InnerHealthError               `json:"innerHealthErrors,omitempty"`
	PossibleCauses               *string                           `json:"possibleCauses,omitempty"`
	RecommendedAction            *string                           `json:"recommendedAction,omitempty"`
	RecoveryProviderErrorMessage *string                           `json:"recoveryProviderErrorMessage,omitempty"`
	SummaryMessage               *string                           `json:"summaryMessage,omitempty"`
}
