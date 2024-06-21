package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleProperties struct {
	Notes                  *string                 `json:"notes,omitempty"`
	ProvisioningState      *ProvisioningState      `json:"provisioningState,omitempty"`
	RecurrencePattern      *RecurrencePattern      `json:"recurrencePattern,omitempty"`
	ResourceOperationError *ResourceOperationError `json:"resourceOperationError,omitempty"`
	StartAt                *string                 `json:"startAt,omitempty"`
	StopAt                 *string                 `json:"stopAt,omitempty"`
	TimeZoneId             *string                 `json:"timeZoneId,omitempty"`
}
