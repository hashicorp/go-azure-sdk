package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureCommunicationServicesUserConversationMemberOperationPredicate struct {
	AzureCommunicationServicesId *string
	DisplayName                  *string
	Id                           *string
	ODataType                    *string
	VisibleHistoryStartDateTime  *string
}

func (p AzureCommunicationServicesUserConversationMemberOperationPredicate) Matches(input AzureCommunicationServicesUserConversationMember) bool {

	if p.AzureCommunicationServicesId != nil && (input.AzureCommunicationServicesId == nil || *p.AzureCommunicationServicesId != *input.AzureCommunicationServicesId) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.VisibleHistoryStartDateTime != nil && (input.VisibleHistoryStartDateTime == nil || *p.VisibleHistoryStartDateTime != *input.VisibleHistoryStartDateTime) {
		return false
	}

	return true
}
