package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateAccountSummary struct {
	CategorySummaries     *[]MacOSSoftwareUpdateCategorySummary `json:"categorySummaries,omitempty"`
	DeviceId              *string                               `json:"deviceId,omitempty"`
	DeviceName            *string                               `json:"deviceName,omitempty"`
	DisplayName           *string                               `json:"displayName,omitempty"`
	FailedUpdateCount     *int64                                `json:"failedUpdateCount,omitempty"`
	Id                    *string                               `json:"id,omitempty"`
	LastUpdatedDateTime   *string                               `json:"lastUpdatedDateTime,omitempty"`
	ODataType             *string                               `json:"@odata.type,omitempty"`
	OsVersion             *string                               `json:"osVersion,omitempty"`
	SuccessfulUpdateCount *int64                                `json:"successfulUpdateCount,omitempty"`
	TotalUpdateCount      *int64                                `json:"totalUpdateCount,omitempty"`
	UserId                *string                               `json:"userId,omitempty"`
	UserPrincipalName     *string                               `json:"userPrincipalName,omitempty"`
}
