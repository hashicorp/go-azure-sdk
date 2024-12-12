package sessionhostmanagements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SessionHostManagementProperties struct {
	ScheduledDateTimeZone string                                `json:"scheduledDateTimeZone"`
	Update                HostPoolUpdateConfigurationProperties `json:"update"`
}
