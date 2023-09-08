package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory struct {
	DeviceId                  *string `json:"deviceId,omitempty"`
	EstimatedRuntimeInMinutes *int64  `json:"estimatedRuntimeInMinutes,omitempty"`
	Id                        *string `json:"id,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
	RuntimeDateTime           *string `json:"runtimeDateTime,omitempty"`
}
