package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatIntelligence struct {
	Confidence        *float64 `json:"confidence,omitempty"`
	ProviderName      *string  `json:"providerName,omitempty"`
	ReportLink        *string  `json:"reportLink,omitempty"`
	ThreatDescription *string  `json:"threatDescription,omitempty"`
	ThreatName        *string  `json:"threatName,omitempty"`
	ThreatType        *string  `json:"threatType,omitempty"`
}
