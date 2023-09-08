package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365GroupsActivityCounts struct {
	ExchangeEmailsReceived *int64  `json:"exchangeEmailsReceived,omitempty"`
	Id                     *string `json:"id,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
	ReportDate             *string `json:"reportDate,omitempty"`
	ReportPeriod           *string `json:"reportPeriod,omitempty"`
	ReportRefreshDate      *string `json:"reportRefreshDate,omitempty"`
	TeamsChannelMessages   *int64  `json:"teamsChannelMessages,omitempty"`
	TeamsMeetingsOrganized *int64  `json:"teamsMeetingsOrganized,omitempty"`
	YammerMessagesLiked    *int64  `json:"yammerMessagesLiked,omitempty"`
	YammerMessagesPosted   *int64  `json:"yammerMessagesPosted,omitempty"`
	YammerMessagesRead     *int64  `json:"yammerMessagesRead,omitempty"`
}
