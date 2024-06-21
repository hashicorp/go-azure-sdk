package recommendedactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationActionProperties struct {
	ActionId           *int64             `json:"actionId,omitempty"`
	AdvisorName        *string            `json:"advisorName,omitempty"`
	CreatedTime        *string            `json:"createdTime,omitempty"`
	Details            *map[string]string `json:"details,omitempty"`
	ExpirationTime     *string            `json:"expirationTime,omitempty"`
	Reason             *string            `json:"reason,omitempty"`
	RecommendationType *string            `json:"recommendationType,omitempty"`
	SessionId          *string            `json:"sessionId,omitempty"`
}
