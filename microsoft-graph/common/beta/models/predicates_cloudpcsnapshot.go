package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSnapshotOperationPredicate struct {
	CloudPCId            *string
	CreatedDateTime      *string
	ExpirationDateTime   *string
	Id                   *string
	LastRestoredDateTime *string
	ODataType            *string
}

func (p CloudPCSnapshotOperationPredicate) Matches(input CloudPCSnapshot) bool {

	if p.CloudPCId != nil && (input.CloudPCId == nil || *p.CloudPCId != *input.CloudPCId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastRestoredDateTime != nil && (input.LastRestoredDateTime == nil || *p.LastRestoredDateTime != *input.LastRestoredDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
