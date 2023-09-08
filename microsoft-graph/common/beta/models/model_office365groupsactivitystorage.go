package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365GroupsActivityStorage struct {
	Id                        *string `json:"id,omitempty"`
	MailboxStorageUsedInBytes *int64  `json:"mailboxStorageUsedInBytes,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
	ReportDate                *string `json:"reportDate,omitempty"`
	ReportPeriod              *string `json:"reportPeriod,omitempty"`
	ReportRefreshDate         *string `json:"reportRefreshDate,omitempty"`
	SiteStorageUsedInBytes    *int64  `json:"siteStorageUsedInBytes,omitempty"`
}
