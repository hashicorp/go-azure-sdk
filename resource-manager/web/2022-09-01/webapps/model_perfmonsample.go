package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PerfMonSample struct {
	InstanceName *string  `json:"instanceName,omitempty"`
	Time         *string  `json:"time,omitempty"`
	Value        *float64 `json:"value,omitempty"`
}
