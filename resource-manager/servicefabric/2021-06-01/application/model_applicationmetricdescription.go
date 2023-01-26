package application

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationMetricDescription struct {
	MaximumCapacity          *int64  `json:"maximumCapacity,omitempty"`
	Name                     *string `json:"name,omitempty"`
	ReservationCapacity      *int64  `json:"reservationCapacity,omitempty"`
	TotalApplicationCapacity *int64  `json:"totalApplicationCapacity,omitempty"`
}
