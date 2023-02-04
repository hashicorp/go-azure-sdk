package jobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ErrorDetails struct {
	AgentErrorCode              *string `json:"agentErrorCode,omitempty"`
	AgentErrorMessage           *string `json:"agentErrorMessage,omitempty"`
	AgentErrorPossibleCauses    *string `json:"agentErrorPossibleCauses,omitempty"`
	AgentErrorRecommendedAction *string `json:"agentErrorRecommendedAction,omitempty"`
	Code                        *string `json:"code,omitempty"`
	IsAgentReportedError        *bool   `json:"isAgentReportedError,omitempty"`
	Message                     *string `json:"message,omitempty"`
	PossibleCauses              *string `json:"possibleCauses,omitempty"`
	RecommendedAction           *string `json:"recommendedAction,omitempty"`
	Severity                    *string `json:"severity,omitempty"`
}
