package meteamwork

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeTeamworkSendActivityNotificationRequest struct {
	ActivityType       *string                `json:"activityType,omitempty"`
	ChainId            *int64                 `json:"chainId,omitempty"`
	PreviewText        *ItemBody              `json:"previewText,omitempty"`
	TemplateParameters *[]KeyValuePair        `json:"templateParameters,omitempty"`
	Topic              *TeamworkActivityTopic `json:"topic,omitempty"`
}
