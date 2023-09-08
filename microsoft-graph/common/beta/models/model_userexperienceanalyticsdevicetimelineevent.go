package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceTimelineEvent struct {
	DeviceId      *string                                               `json:"deviceId,omitempty"`
	EventDateTime *string                                               `json:"eventDateTime,omitempty"`
	EventDetails  *string                                               `json:"eventDetails,omitempty"`
	EventLevel    *UserExperienceAnalyticsDeviceTimelineEventEventLevel `json:"eventLevel,omitempty"`
	EventName     *string                                               `json:"eventName,omitempty"`
	EventSource   *string                                               `json:"eventSource,omitempty"`
	Id            *string                                               `json:"id,omitempty"`
	ODataType     *string                                               `json:"@odata.type,omitempty"`
}
