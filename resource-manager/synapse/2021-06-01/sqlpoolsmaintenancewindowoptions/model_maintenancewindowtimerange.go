package sqlpoolsmaintenancewindowoptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaintenanceWindowTimeRange struct {
	DayOfWeek *DayOfWeek `json:"dayOfWeek,omitempty"`
	Duration  *string    `json:"duration,omitempty"`
	StartTime *string    `json:"startTime,omitempty"`
}
