package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptRunSummary struct {
	ErrorDeviceCount   *int64  `json:"errorDeviceCount,omitempty"`
	ErrorUserCount     *int64  `json:"errorUserCount,omitempty"`
	Id                 *string `json:"id,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	SuccessDeviceCount *int64  `json:"successDeviceCount,omitempty"`
	SuccessUserCount   *int64  `json:"successUserCount,omitempty"`
}
