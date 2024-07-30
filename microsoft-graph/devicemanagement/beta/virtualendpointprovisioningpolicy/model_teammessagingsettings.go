package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamMessagingSettings struct {
	AllowChannelMentions     *bool   `json:"allowChannelMentions,omitempty"`
	AllowOwnerDeleteMessages *bool   `json:"allowOwnerDeleteMessages,omitempty"`
	AllowTeamMentions        *bool   `json:"allowTeamMentions,omitempty"`
	AllowUserDeleteMessages  *bool   `json:"allowUserDeleteMessages,omitempty"`
	AllowUserEditMessages    *bool   `json:"allowUserEditMessages,omitempty"`
	ODataType                *string `json:"@odata.type,omitempty"`
}
