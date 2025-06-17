package chat

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendChatActivityNotificationRequest struct {
	ActivityType       nullable.Type[string]               `json:"activityType,omitempty"`
	ChainId            nullable.Type[int64]                `json:"chainId,omitempty"`
	IconId             nullable.Type[string]               `json:"iconId,omitempty"`
	PreviewText        *beta.ItemBody                      `json:"previewText,omitempty"`
	Recipient          *beta.TeamworkNotificationRecipient `json:"recipient,omitempty"`
	TeamsAppId         nullable.Type[string]               `json:"teamsAppId,omitempty"`
	TemplateParameters *[]beta.KeyValuePair                `json:"templateParameters,omitempty"`
	Topic              *beta.TeamworkActivityTopic         `json:"topic,omitempty"`
}
