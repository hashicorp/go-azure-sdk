package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVppTokenTroubleshootingEventOperationPredicate struct {
	CorrelationId *string
	EventDateTime *string
	EventName     *string
	Id            *string
	ODataType     *string
	TokenId       *string
}

func (p AppleVppTokenTroubleshootingEventOperationPredicate) Matches(input AppleVppTokenTroubleshootingEvent) bool {

	if p.CorrelationId != nil && (input.CorrelationId == nil || *p.CorrelationId != *input.CorrelationId) {
		return false
	}

	if p.EventDateTime != nil && (input.EventDateTime == nil || *p.EventDateTime != *input.EventDateTime) {
		return false
	}

	if p.EventName != nil && (input.EventName == nil || *p.EventName != *input.EventName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TokenId != nil && (input.TokenId == nil || *p.TokenId != *input.TokenId) {
		return false
	}

	return true
}
