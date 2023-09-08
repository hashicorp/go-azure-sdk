package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorStatusDetailsOperationPredicate struct {
	ConnectorInstanceId *string
	EventDateTime       *string
	ODataType           *string
}

func (p ConnectorStatusDetailsOperationPredicate) Matches(input ConnectorStatusDetails) bool {

	if p.ConnectorInstanceId != nil && (input.ConnectorInstanceId == nil || *p.ConnectorInstanceId != *input.ConnectorInstanceId) {
		return false
	}

	if p.EventDateTime != nil && (input.EventDateTime == nil || *p.EventDateTime != *input.EventDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
