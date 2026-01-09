package team

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendTeamActivityNotificationRequest struct {
	ActivityType       nullable.Type[string]               `json:"activityType,omitempty"`
	ChainId            nullable.Type[int64]                `json:"chainId,omitempty"`
	IconId             nullable.Type[string]               `json:"iconId,omitempty"`
	PreviewText        *beta.ItemBody                      `json:"previewText,omitempty"`
	Recipient          *beta.TeamworkNotificationRecipient `json:"recipient,omitempty"`
	TeamsAppId         nullable.Type[string]               `json:"teamsAppId,omitempty"`
	TemplateParameters *[]beta.KeyValuePair                `json:"templateParameters,omitempty"`
	Topic              *beta.TeamworkActivityTopic         `json:"topic,omitempty"`
}
