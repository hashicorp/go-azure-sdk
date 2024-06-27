package reservations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationUtilizationAggregates struct {
	Grain     *float64 `json:"grain,omitempty"`
	GrainUnit *string  `json:"grainUnit,omitempty"`
	Value     *float64 `json:"value,omitempty"`
	ValueUnit *string  `json:"valueUnit,omitempty"`
}
