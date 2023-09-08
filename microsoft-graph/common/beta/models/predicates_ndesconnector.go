package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NdesConnectorOperationPredicate struct {
	ConnectorVersion       *string
	DisplayName            *string
	EnrolledDateTime       *string
	Id                     *string
	LastConnectionDateTime *string
	MachineName            *string
	ODataType              *string
}

func (p NdesConnectorOperationPredicate) Matches(input NdesConnector) bool {

	if p.ConnectorVersion != nil && (input.ConnectorVersion == nil || *p.ConnectorVersion != *input.ConnectorVersion) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EnrolledDateTime != nil && (input.EnrolledDateTime == nil || *p.EnrolledDateTime != *input.EnrolledDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastConnectionDateTime != nil && (input.LastConnectionDateTime == nil || *p.LastConnectionDateTime != *input.LastConnectionDateTime) {
		return false
	}

	if p.MachineName != nil && (input.MachineName == nil || *p.MachineName != *input.MachineName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
