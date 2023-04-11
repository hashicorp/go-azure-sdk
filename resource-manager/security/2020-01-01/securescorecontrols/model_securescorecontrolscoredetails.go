package securescorecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoreControlScoreDetails struct {
	Definition                 *SecureScoreControlDefinitionItem `json:"definition,omitempty"`
	DisplayName                *string                           `json:"displayName,omitempty"`
	HealthyResourceCount       *int64                            `json:"healthyResourceCount,omitempty"`
	NotApplicableResourceCount *int64                            `json:"notApplicableResourceCount,omitempty"`
	Score                      *ScoreDetails                     `json:"score,omitempty"`
	UnhealthyResourceCount     *int64                            `json:"unhealthyResourceCount,omitempty"`
	Weight                     *int64                            `json:"weight,omitempty"`
}
