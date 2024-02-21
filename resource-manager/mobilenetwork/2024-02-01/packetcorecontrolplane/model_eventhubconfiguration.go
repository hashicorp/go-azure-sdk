package packetcorecontrolplane

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventHubConfiguration struct {
	Id                string `json:"id"`
	ReportingInterval *int64 `json:"reportingInterval,omitempty"`
}
