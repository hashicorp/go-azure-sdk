package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingParticipantInfoOperationPredicate struct {
	ODataType *string
	Upn       *string
}

func (p MeetingParticipantInfoOperationPredicate) Matches(input MeetingParticipantInfo) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Upn != nil && (input.Upn == nil || *p.Upn != *input.Upn) {
		return false
	}

	return true
}
