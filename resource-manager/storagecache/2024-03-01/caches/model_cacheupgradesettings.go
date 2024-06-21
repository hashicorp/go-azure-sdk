package caches

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CacheUpgradeSettings struct {
	ScheduledTime          *string `json:"scheduledTime,omitempty"`
	UpgradeScheduleEnabled *bool   `json:"upgradeScheduleEnabled,omitempty"`
}
