package replicationeligibilityresults

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationEligibilityResultsErrorInfo struct {
	Code              *string `json:"code,omitempty"`
	Message           *string `json:"message,omitempty"`
	PossibleCauses    *string `json:"possibleCauses,omitempty"`
	RecommendedAction *string `json:"recommendedAction,omitempty"`
	Status            *string `json:"status,omitempty"`
}
