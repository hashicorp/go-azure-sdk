package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectOnlineMeetingAction struct {
	AllowedForwarders        *ProtectOnlineMeetingActionAllowedForwarders `json:"allowedForwarders,omitempty"`
	AllowedPresenters        *ProtectOnlineMeetingActionAllowedPresenters `json:"allowedPresenters,omitempty"`
	IsCopyToClipboardEnabled *bool                                        `json:"isCopyToClipboardEnabled,omitempty"`
	IsLobbyEnabled           *bool                                        `json:"isLobbyEnabled,omitempty"`
	LobbyBypassSettings      *LobbyBypassSettings                         `json:"lobbyBypassSettings,omitempty"`
	Name                     *string                                      `json:"name,omitempty"`
	ODataType                *string                                      `json:"@odata.type,omitempty"`
}
