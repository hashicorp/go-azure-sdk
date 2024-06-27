package reservation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationExtendedStatusInfo struct {
	Message    *string                             `json:"message,omitempty"`
	Properties *ExtendedStatusDefinitionProperties `json:"properties,omitempty"`
	StatusCode *ReservationStatusCode              `json:"statusCode,omitempty"`
}
