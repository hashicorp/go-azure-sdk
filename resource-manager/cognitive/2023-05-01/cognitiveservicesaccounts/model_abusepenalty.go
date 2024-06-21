package cognitiveservicesaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AbusePenalty struct {
	Action              *AbusePenaltyAction `json:"action,omitempty"`
	Expiration          *string             `json:"expiration,omitempty"`
	RateLimitPercentage *float64            `json:"rateLimitPercentage,omitempty"`
}
