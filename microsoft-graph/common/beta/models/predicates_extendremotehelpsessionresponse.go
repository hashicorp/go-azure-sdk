package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtendRemoteHelpSessionResponseOperationPredicate struct {
	AcsHelperUserToken        *string
	ODataType                 *string
	PubSubHelperAccessUri     *string
	SessionExpirationDateTime *string
	SessionKey                *string
}

func (p ExtendRemoteHelpSessionResponseOperationPredicate) Matches(input ExtendRemoteHelpSessionResponse) bool {

	if p.AcsHelperUserToken != nil && (input.AcsHelperUserToken == nil || *p.AcsHelperUserToken != *input.AcsHelperUserToken) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PubSubHelperAccessUri != nil && (input.PubSubHelperAccessUri == nil || *p.PubSubHelperAccessUri != *input.PubSubHelperAccessUri) {
		return false
	}

	if p.SessionExpirationDateTime != nil && (input.SessionExpirationDateTime == nil || *p.SessionExpirationDateTime != *input.SessionExpirationDateTime) {
		return false
	}

	if p.SessionKey != nil && (input.SessionKey == nil || *p.SessionKey != *input.SessionKey) {
		return false
	}

	return true
}
