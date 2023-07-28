package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringDataSegment struct {
	Feature *string   `json:"feature,omitempty"`
	Values  *[]string `json:"values,omitempty"`
}
