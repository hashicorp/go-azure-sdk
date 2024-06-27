package reservation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationPropertyUtilization struct {
	Aggregates *[]ReservationUtilizationAggregates `json:"aggregates,omitempty"`
	Trend      *string                             `json:"trend,omitempty"`
}
