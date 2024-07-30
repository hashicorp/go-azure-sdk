package me

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Mention struct {
	Application           *string       `json:"application,omitempty"`
	ClientReference       *string       `json:"clientReference,omitempty"`
	CreatedBy             *EmailAddress `json:"createdBy,omitempty"`
	CreatedDateTime       *string       `json:"createdDateTime,omitempty"`
	DeepLink              *string       `json:"deepLink,omitempty"`
	Id                    *string       `json:"id,omitempty"`
	MentionText           *string       `json:"mentionText,omitempty"`
	Mentioned             *EmailAddress `json:"mentioned,omitempty"`
	ODataType             *string       `json:"@odata.type,omitempty"`
	ServerCreatedDateTime *string       `json:"serverCreatedDateTime,omitempty"`
}
