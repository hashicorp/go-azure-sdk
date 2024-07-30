package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthDevicePerformanceDetails struct {
	AppDisplayName    *string `json:"appDisplayName,omitempty"`
	AppPublisher      *string `json:"appPublisher,omitempty"`
	AppVersion        *string `json:"appVersion,omitempty"`
	DeviceDisplayName *string `json:"deviceDisplayName,omitempty"`
	DeviceId          *string `json:"deviceId,omitempty"`
	EventDateTime     *string `json:"eventDateTime,omitempty"`
	EventType         *string `json:"eventType,omitempty"`
	Id                *string `json:"id,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
}
