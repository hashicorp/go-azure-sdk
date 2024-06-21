package networkconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthCheckStatusDetailsProperties struct {
	EndDateTime   *string        `json:"endDateTime,omitempty"`
	HealthChecks  *[]HealthCheck `json:"healthChecks,omitempty"`
	StartDateTime *string        `json:"startDateTime,omitempty"`
}
