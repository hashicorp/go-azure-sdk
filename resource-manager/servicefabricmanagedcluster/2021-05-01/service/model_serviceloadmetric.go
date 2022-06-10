package service

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceLoadMetric struct {
	DefaultLoad          *int64                   `json:"defaultLoad,omitempty"`
	Name                 string                   `json:"name"`
	PrimaryDefaultLoad   *int64                   `json:"primaryDefaultLoad,omitempty"`
	SecondaryDefaultLoad *int64                   `json:"secondaryDefaultLoad,omitempty"`
	Weight               *ServiceLoadMetricWeight `json:"weight,omitempty"`
}
