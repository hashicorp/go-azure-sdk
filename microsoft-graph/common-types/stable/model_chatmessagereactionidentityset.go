package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageReactionIdentitySet struct {
	Application *Identity `json:"application,omitempty"`
	Device      *Identity `json:"device,omitempty"`
	ODataType   *string   `json:"@odata.type,omitempty"`
	User        *Identity `json:"user,omitempty"`
}
