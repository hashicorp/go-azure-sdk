package reservedinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportURL struct {
	ReportUrl  *ReservationReportSchema `json:"reportUrl,omitempty"`
	ValidUntil *string                  `json:"validUntil,omitempty"`
}
