package pricesheetresults

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SavingsPlan struct {
	EffectivePrice *float64 `json:"effectivePrice,omitempty"`
	MarketPrice    *float64 `json:"marketPrice,omitempty"`
	Term           *string  `json:"term,omitempty"`
}
