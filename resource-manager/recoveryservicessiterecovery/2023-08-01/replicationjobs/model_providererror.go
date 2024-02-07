package replicationjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderError struct {
	ErrorCode         *int64  `json:"errorCode,omitempty"`
	ErrorId           *string `json:"errorId,omitempty"`
	ErrorMessage      *string `json:"errorMessage,omitempty"`
	PossibleCauses    *string `json:"possibleCauses,omitempty"`
	RecommendedAction *string `json:"recommendedAction,omitempty"`
}
