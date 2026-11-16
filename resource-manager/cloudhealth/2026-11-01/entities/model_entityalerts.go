package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityAlerts struct {
	Degraded  *AlertConfiguration `json:"degraded,omitempty"`
	Unhealthy *AlertConfiguration `json:"unhealthy,omitempty"`
}
