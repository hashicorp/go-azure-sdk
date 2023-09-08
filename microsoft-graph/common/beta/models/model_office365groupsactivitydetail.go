package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365GroupsActivityDetail struct {
	ExchangeMailboxStorageUsedInBytes *int64  `json:"exchangeMailboxStorageUsedInBytes,omitempty"`
	ExchangeMailboxTotalItemCount     *int64  `json:"exchangeMailboxTotalItemCount,omitempty"`
	ExchangeReceivedEmailCount        *int64  `json:"exchangeReceivedEmailCount,omitempty"`
	ExternalMemberCount               *int64  `json:"externalMemberCount,omitempty"`
	GroupDisplayName                  *string `json:"groupDisplayName,omitempty"`
	GroupId                           *string `json:"groupId,omitempty"`
	GroupType                         *string `json:"groupType,omitempty"`
	Id                                *string `json:"id,omitempty"`
	IsDeleted                         *bool   `json:"isDeleted,omitempty"`
	LastActivityDate                  *string `json:"lastActivityDate,omitempty"`
	MemberCount                       *int64  `json:"memberCount,omitempty"`
	ODataType                         *string `json:"@odata.type,omitempty"`
	OwnerPrincipalName                *string `json:"ownerPrincipalName,omitempty"`
	ReportPeriod                      *string `json:"reportPeriod,omitempty"`
	ReportRefreshDate                 *string `json:"reportRefreshDate,omitempty"`
	SharePointActiveFileCount         *int64  `json:"sharePointActiveFileCount,omitempty"`
	SharePointSiteStorageUsedInBytes  *int64  `json:"sharePointSiteStorageUsedInBytes,omitempty"`
	SharePointTotalFileCount          *int64  `json:"sharePointTotalFileCount,omitempty"`
	TeamsChannelMessagesCount         *int64  `json:"teamsChannelMessagesCount,omitempty"`
	TeamsMeetingsOrganizedCount       *int64  `json:"teamsMeetingsOrganizedCount,omitempty"`
	YammerLikedMessageCount           *int64  `json:"yammerLikedMessageCount,omitempty"`
	YammerPostedMessageCount          *int64  `json:"yammerPostedMessageCount,omitempty"`
	YammerReadMessageCount            *int64  `json:"yammerReadMessageCount,omitempty"`
}
