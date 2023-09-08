package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthAppImpact struct {
	ActiveDevices   *int64  `json:"activeDevices,omitempty"`
	AppDisplayName  *string `json:"appDisplayName,omitempty"`
	AppName         *string `json:"appName,omitempty"`
	AppPublisher    *string `json:"appPublisher,omitempty"`
	Id              *string `json:"id,omitempty"`
	IsForegroundApp *bool   `json:"isForegroundApp,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
}
