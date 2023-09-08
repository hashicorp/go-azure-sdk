package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityIssue struct {
	DeviceId          *string `json:"deviceId,omitempty"`
	ErrorCode         *string `json:"errorCode,omitempty"`
	ErrorDateTime     *string `json:"errorDateTime,omitempty"`
	ErrorDescription  *string `json:"errorDescription,omitempty"`
	Id                *string `json:"id,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	RecommendedAction *string `json:"recommendedAction,omitempty"`
	UserId            *string `json:"userId,omitempty"`
}
