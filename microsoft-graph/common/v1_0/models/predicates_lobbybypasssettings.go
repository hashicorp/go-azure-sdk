package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LobbyBypassSettingsOperationPredicate struct {
	IsDialInBypassEnabled *bool
	ODataType             *string
}

func (p LobbyBypassSettingsOperationPredicate) Matches(input LobbyBypassSettings) bool {

	if p.IsDialInBypassEnabled != nil && (input.IsDialInBypassEnabled == nil || *p.IsDialInBypassEnabled != *input.IsDialInBypassEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
