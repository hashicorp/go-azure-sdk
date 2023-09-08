package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365GroupsActivityCountsOperationPredicate struct {
	ExchangeEmailsReceived *int64
	Id                     *string
	ODataType              *string
	ReportDate             *string
	ReportPeriod           *string
	ReportRefreshDate      *string
	TeamsChannelMessages   *int64
	TeamsMeetingsOrganized *int64
	YammerMessagesLiked    *int64
	YammerMessagesPosted   *int64
	YammerMessagesRead     *int64
}

func (p Office365GroupsActivityCountsOperationPredicate) Matches(input Office365GroupsActivityCounts) bool {

	if p.ExchangeEmailsReceived != nil && (input.ExchangeEmailsReceived == nil || *p.ExchangeEmailsReceived != *input.ExchangeEmailsReceived) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReportDate != nil && (input.ReportDate == nil || *p.ReportDate != *input.ReportDate) {
		return false
	}

	if p.ReportPeriod != nil && (input.ReportPeriod == nil || *p.ReportPeriod != *input.ReportPeriod) {
		return false
	}

	if p.ReportRefreshDate != nil && (input.ReportRefreshDate == nil || *p.ReportRefreshDate != *input.ReportRefreshDate) {
		return false
	}

	if p.TeamsChannelMessages != nil && (input.TeamsChannelMessages == nil || *p.TeamsChannelMessages != *input.TeamsChannelMessages) {
		return false
	}

	if p.TeamsMeetingsOrganized != nil && (input.TeamsMeetingsOrganized == nil || *p.TeamsMeetingsOrganized != *input.TeamsMeetingsOrganized) {
		return false
	}

	if p.YammerMessagesLiked != nil && (input.YammerMessagesLiked == nil || *p.YammerMessagesLiked != *input.YammerMessagesLiked) {
		return false
	}

	if p.YammerMessagesPosted != nil && (input.YammerMessagesPosted == nil || *p.YammerMessagesPosted != *input.YammerMessagesPosted) {
		return false
	}

	if p.YammerMessagesRead != nil && (input.YammerMessagesRead == nil || *p.YammerMessagesRead != *input.YammerMessagesRead) {
		return false
	}

	return true
}
