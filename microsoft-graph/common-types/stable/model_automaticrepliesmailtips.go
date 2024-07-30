package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticRepliesMailTips struct {
	Message            *string           `json:"message,omitempty"`
	MessageLanguage    *LocaleInfo       `json:"messageLanguage,omitempty"`
	ODataType          *string           `json:"@odata.type,omitempty"`
	ScheduledEndTime   *DateTimeTimeZone `json:"scheduledEndTime,omitempty"`
	ScheduledStartTime *DateTimeTimeZone `json:"scheduledStartTime,omitempty"`
}
