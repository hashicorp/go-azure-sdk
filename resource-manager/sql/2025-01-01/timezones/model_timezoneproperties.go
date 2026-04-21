package timezones

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeZoneProperties struct {
	DisplayName *string `json:"displayName,omitempty"`
	TimeZoneId  *string `json:"timeZoneId,omitempty"`
}
