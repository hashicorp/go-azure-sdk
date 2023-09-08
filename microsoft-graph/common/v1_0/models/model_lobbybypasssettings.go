package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LobbyBypassSettings struct {
	IsDialInBypassEnabled *bool                     `json:"isDialInBypassEnabled,omitempty"`
	ODataType             *string                   `json:"@odata.type,omitempty"`
	Scope                 *LobbyBypassSettingsScope `json:"scope,omitempty"`
}
