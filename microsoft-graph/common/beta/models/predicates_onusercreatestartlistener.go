package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnUserCreateStartListenerOperationPredicate struct {
	AuthenticationEventsFlowId *string
	Id                         *string
	ODataType                  *string
	Priority                   *int64
}

func (p OnUserCreateStartListenerOperationPredicate) Matches(input OnUserCreateStartListener) bool {

	if p.AuthenticationEventsFlowId != nil && (input.AuthenticationEventsFlowId == nil || *p.AuthenticationEventsFlowId != *input.AuthenticationEventsFlowId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Priority != nil && (input.Priority == nil || *p.Priority != *input.Priority) {
		return false
	}

	return true
}
