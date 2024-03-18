package replicationmigrationitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplianceResourceDetails struct {
	Capacity           *int64   `json:"capacity,omitempty"`
	ProcessUtilization *float64 `json:"processUtilization,omitempty"`
	Status             *string  `json:"status,omitempty"`
	TotalUtilization   *float64 `json:"totalUtilization,omitempty"`
}
