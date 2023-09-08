package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalConnectionOperationPredicate struct {
	ConnectorId        *string
	Description        *string
	Id                 *string
	IngestedItemsCount *int64
	Name               *string
	ODataType          *string
}

func (p ExternalConnectorsExternalConnectionOperationPredicate) Matches(input ExternalConnectorsExternalConnection) bool {

	if p.ConnectorId != nil && (input.ConnectorId == nil || *p.ConnectorId != *input.ConnectorId) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IngestedItemsCount != nil && (input.IngestedItemsCount == nil || *p.IngestedItemsCount != *input.IngestedItemsCount) {
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
