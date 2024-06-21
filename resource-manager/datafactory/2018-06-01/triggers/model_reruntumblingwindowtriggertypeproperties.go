package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RerunTumblingWindowTriggerTypeProperties struct {
	ParentTrigger      interface{} `json:"parentTrigger"`
	RequestedEndTime   string      `json:"requestedEndTime"`
	RequestedStartTime string      `json:"requestedStartTime"`
	RerunConcurrency   int64       `json:"rerunConcurrency"`
}
