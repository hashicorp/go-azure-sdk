package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveUsersBreakdownMetric struct {
	AppId     *string `json:"appId,omitempty"`
	AppName   *string `json:"appName,omitempty"`
	Count     *int64  `json:"count,omitempty"`
	FactDate  *string `json:"factDate,omitempty"`
	Id        *string `json:"id,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Os        *string `json:"os,omitempty"`
}
