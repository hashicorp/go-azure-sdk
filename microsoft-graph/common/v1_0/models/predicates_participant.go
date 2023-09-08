package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParticipantOperationPredicate struct {
	Id        *string
	IsInLobby *bool
	IsMuted   *bool
	Metadata  *string
	ODataType *string
}

func (p ParticipantOperationPredicate) Matches(input Participant) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsInLobby != nil && (input.IsInLobby == nil || *p.IsInLobby != *input.IsInLobby) {
		return false
	}

	if p.IsMuted != nil && (input.IsMuted == nil || *p.IsMuted != *input.IsMuted) {
		return false
	}

	if p.Metadata != nil && (input.Metadata == nil || *p.Metadata != *input.Metadata) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
