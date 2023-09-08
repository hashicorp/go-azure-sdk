package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectOnlineMeetingActionOperationPredicate struct {
	IsCopyToClipboardEnabled *bool
	IsLobbyEnabled           *bool
	Name                     *string
	ODataType                *string
}

func (p ProtectOnlineMeetingActionOperationPredicate) Matches(input ProtectOnlineMeetingAction) bool {

	if p.IsCopyToClipboardEnabled != nil && (input.IsCopyToClipboardEnabled == nil || *p.IsCopyToClipboardEnabled != *input.IsCopyToClipboardEnabled) {
		return false
	}

	if p.IsLobbyEnabled != nil && (input.IsLobbyEnabled == nil || *p.IsLobbyEnabled != *input.IsLobbyEnabled) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
