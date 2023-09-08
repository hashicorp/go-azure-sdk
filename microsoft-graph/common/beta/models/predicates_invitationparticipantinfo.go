package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationParticipantInfoOperationPredicate struct {
	Hidden                             *bool
	ODataType                          *string
	ParticipantId                      *string
	RemoveFromDefaultAudioRoutingGroup *bool
	ReplacesCallId                     *string
}

func (p InvitationParticipantInfoOperationPredicate) Matches(input InvitationParticipantInfo) bool {

	if p.Hidden != nil && (input.Hidden == nil || *p.Hidden != *input.Hidden) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ParticipantId != nil && (input.ParticipantId == nil || *p.ParticipantId != *input.ParticipantId) {
		return false
	}

	if p.RemoveFromDefaultAudioRoutingGroup != nil && (input.RemoveFromDefaultAudioRoutingGroup == nil || *p.RemoveFromDefaultAudioRoutingGroup != *input.RemoveFromDefaultAudioRoutingGroup) {
		return false
	}

	if p.ReplacesCallId != nil && (input.ReplacesCallId == nil || *p.ReplacesCallId != *input.ReplacesCallId) {
		return false
	}

	return true
}
