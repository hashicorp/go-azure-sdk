package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationUsageDetails struct {
	Charges    *[]float64 `json:"charges,omitempty"`
	UsageGrain *Grain     `json:"usageGrain,omitempty"`
}
