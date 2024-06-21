package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PeriodicTimerSourceInfo struct {
	Schedule  string  `json:"schedule"`
	StartTime string  `json:"startTime"`
	Topic     *string `json:"topic,omitempty"`
}
