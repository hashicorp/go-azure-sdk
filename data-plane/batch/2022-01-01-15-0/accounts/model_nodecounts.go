package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeCounts struct {
	Creating            int64 `json:"creating"`
	Idle                int64 `json:"idle"`
	LeavingPool         int64 `json:"leavingPool"`
	Offline             int64 `json:"offline"`
	Preempted           int64 `json:"preempted"`
	Rebooting           int64 `json:"rebooting"`
	Reimaging           int64 `json:"reimaging"`
	Running             int64 `json:"running"`
	StartTaskFailed     int64 `json:"startTaskFailed"`
	Starting            int64 `json:"starting"`
	Total               int64 `json:"total"`
	Unknown             int64 `json:"unknown"`
	Unusable            int64 `json:"unusable"`
	WaitingForStartTask int64 `json:"waitingForStartTask"`
}
