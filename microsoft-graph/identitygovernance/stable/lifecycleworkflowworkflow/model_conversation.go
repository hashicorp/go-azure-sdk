package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Conversation struct {
	HasAttachments        *bool                 `json:"hasAttachments,omitempty"`
	Id                    *string               `json:"id,omitempty"`
	LastDeliveredDateTime *string               `json:"lastDeliveredDateTime,omitempty"`
	ODataType             *string               `json:"@odata.type,omitempty"`
	Preview               *string               `json:"preview,omitempty"`
	Threads               *[]ConversationThread `json:"threads,omitempty"`
	Topic                 *string               `json:"topic,omitempty"`
	UniqueSenders         *[]string             `json:"uniqueSenders,omitempty"`
}
