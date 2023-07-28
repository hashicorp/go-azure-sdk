package modelversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineInferenceConfiguration struct {
	Configurations *map[string]string `json:"configurations,omitempty"`
	EntryScript    *string            `json:"entryScript,omitempty"`
	LivenessRoute  *Route             `json:"livenessRoute,omitempty"`
	ReadinessRoute *Route             `json:"readinessRoute,omitempty"`
	ScoringRoute   *Route             `json:"scoringRoute,omitempty"`
}
