package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365GroupsActivityDetailOperationPredicate struct {
	ExchangeMailboxStorageUsedInBytes *int64
	ExchangeMailboxTotalItemCount     *int64
	ExchangeReceivedEmailCount        *int64
	ExternalMemberCount               *int64
	GroupDisplayName                  *string
	GroupId                           *string
	GroupType                         *string
	Id                                *string
	IsDeleted                         *bool
	LastActivityDate                  *string
	MemberCount                       *int64
	ODataType                         *string
	OwnerPrincipalName                *string
	ReportPeriod                      *string
	ReportRefreshDate                 *string
	SharePointActiveFileCount         *int64
	SharePointSiteStorageUsedInBytes  *int64
	SharePointTotalFileCount          *int64
	TeamsChannelMessagesCount         *int64
	TeamsMeetingsOrganizedCount       *int64
	YammerLikedMessageCount           *int64
	YammerPostedMessageCount          *int64
	YammerReadMessageCount            *int64
}

func (p Office365GroupsActivityDetailOperationPredicate) Matches(input Office365GroupsActivityDetail) bool {

	if p.ExchangeMailboxStorageUsedInBytes != nil && (input.ExchangeMailboxStorageUsedInBytes == nil || *p.ExchangeMailboxStorageUsedInBytes != *input.ExchangeMailboxStorageUsedInBytes) {
		return false
	}

	if p.ExchangeMailboxTotalItemCount != nil && (input.ExchangeMailboxTotalItemCount == nil || *p.ExchangeMailboxTotalItemCount != *input.ExchangeMailboxTotalItemCount) {
		return false
	}

	if p.ExchangeReceivedEmailCount != nil && (input.ExchangeReceivedEmailCount == nil || *p.ExchangeReceivedEmailCount != *input.ExchangeReceivedEmailCount) {
		return false
	}

	if p.ExternalMemberCount != nil && (input.ExternalMemberCount == nil || *p.ExternalMemberCount != *input.ExternalMemberCount) {
		return false
	}

	if p.GroupDisplayName != nil && (input.GroupDisplayName == nil || *p.GroupDisplayName != *input.GroupDisplayName) {
		return false
	}

	if p.GroupId != nil && (input.GroupId == nil || *p.GroupId != *input.GroupId) {
		return false
	}

	if p.GroupType != nil && (input.GroupType == nil || *p.GroupType != *input.GroupType) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDeleted != nil && (input.IsDeleted == nil || *p.IsDeleted != *input.IsDeleted) {
		return false
	}

	if p.LastActivityDate != nil && (input.LastActivityDate == nil || *p.LastActivityDate != *input.LastActivityDate) {
		return false
	}

	if p.MemberCount != nil && (input.MemberCount == nil || *p.MemberCount != *input.MemberCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OwnerPrincipalName != nil && (input.OwnerPrincipalName == nil || *p.OwnerPrincipalName != *input.OwnerPrincipalName) {
		return false
	}

	if p.ReportPeriod != nil && (input.ReportPeriod == nil || *p.ReportPeriod != *input.ReportPeriod) {
		return false
	}

	if p.ReportRefreshDate != nil && (input.ReportRefreshDate == nil || *p.ReportRefreshDate != *input.ReportRefreshDate) {
		return false
	}

	if p.SharePointActiveFileCount != nil && (input.SharePointActiveFileCount == nil || *p.SharePointActiveFileCount != *input.SharePointActiveFileCount) {
		return false
	}

	if p.SharePointSiteStorageUsedInBytes != nil && (input.SharePointSiteStorageUsedInBytes == nil || *p.SharePointSiteStorageUsedInBytes != *input.SharePointSiteStorageUsedInBytes) {
		return false
	}

	if p.SharePointTotalFileCount != nil && (input.SharePointTotalFileCount == nil || *p.SharePointTotalFileCount != *input.SharePointTotalFileCount) {
		return false
	}

	if p.TeamsChannelMessagesCount != nil && (input.TeamsChannelMessagesCount == nil || *p.TeamsChannelMessagesCount != *input.TeamsChannelMessagesCount) {
		return false
	}

	if p.TeamsMeetingsOrganizedCount != nil && (input.TeamsMeetingsOrganizedCount == nil || *p.TeamsMeetingsOrganizedCount != *input.TeamsMeetingsOrganizedCount) {
		return false
	}

	if p.YammerLikedMessageCount != nil && (input.YammerLikedMessageCount == nil || *p.YammerLikedMessageCount != *input.YammerLikedMessageCount) {
		return false
	}

	if p.YammerPostedMessageCount != nil && (input.YammerPostedMessageCount == nil || *p.YammerPostedMessageCount != *input.YammerPostedMessageCount) {
		return false
	}

	if p.YammerReadMessageCount != nil && (input.YammerReadMessageCount == nil || *p.YammerReadMessageCount != *input.YammerReadMessageCount) {
		return false
	}

	return true
}
