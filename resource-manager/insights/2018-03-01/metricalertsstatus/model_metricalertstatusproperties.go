package metricalertsstatus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricAlertStatusProperties struct {
	Dimensions *map[string]string `json:"dimensions,omitempty"`
	Status     *string            `json:"status,omitempty"`
	Timestamp  *string            `json:"timestamp,omitempty"`
}
