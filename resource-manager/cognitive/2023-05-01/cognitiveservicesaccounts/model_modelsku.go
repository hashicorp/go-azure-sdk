package cognitiveservicesaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModelSku struct {
	Capacity        *CapacityConfig  `json:"capacity,omitempty"`
	DeprecationDate *string          `json:"deprecationDate,omitempty"`
	Name            *string          `json:"name,omitempty"`
	RateLimits      *[]CallRateLimit `json:"rateLimits,omitempty"`
	UsageName       *string          `json:"usageName,omitempty"`
}
