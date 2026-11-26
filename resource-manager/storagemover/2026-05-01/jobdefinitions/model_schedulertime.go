package jobdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchedulerTime struct {
	Hour   *int64  `json:"hour,omitempty"`
	Minute *Minute `json:"minute,omitempty"`
}
