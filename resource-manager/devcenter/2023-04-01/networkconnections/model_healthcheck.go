package networkconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthCheck struct {
	AdditionalDetails *string            `json:"additionalDetails,omitempty"`
	DisplayName       *string            `json:"displayName,omitempty"`
	EndDateTime       *string            `json:"endDateTime,omitempty"`
	ErrorType         *string            `json:"errorType,omitempty"`
	RecommendedAction *string            `json:"recommendedAction,omitempty"`
	StartDateTime     *string            `json:"startDateTime,omitempty"`
	Status            *HealthCheckStatus `json:"status,omitempty"`
}
