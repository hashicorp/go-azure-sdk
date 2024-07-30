package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatViewpoint struct {
	IsHidden                *bool   `json:"isHidden,omitempty"`
	LastMessageReadDateTime *string `json:"lastMessageReadDateTime,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
}
