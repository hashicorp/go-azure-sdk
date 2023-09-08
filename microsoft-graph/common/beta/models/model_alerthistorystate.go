package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertHistoryState struct {
	AppId           *string                    `json:"appId,omitempty"`
	AssignedTo      *string                    `json:"assignedTo,omitempty"`
	Comments        *[]string                  `json:"comments,omitempty"`
	Feedback        *AlertHistoryStateFeedback `json:"feedback,omitempty"`
	ODataType       *string                    `json:"@odata.type,omitempty"`
	Status          *AlertHistoryStateStatus   `json:"status,omitempty"`
	UpdatedDateTime *string                    `json:"updatedDateTime,omitempty"`
	User            *string                    `json:"user,omitempty"`
}
