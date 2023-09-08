package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardWebPartOperationPredicate struct {
	ContainerTextWebPartId *string
	Id                     *string
	ODataType              *string
	WebPartType            *string
}

func (p StandardWebPartOperationPredicate) Matches(input StandardWebPart) bool {

	if p.ContainerTextWebPartId != nil && (input.ContainerTextWebPartId == nil || *p.ContainerTextWebPartId != *input.ContainerTextWebPartId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.WebPartType != nil && (input.WebPartType == nil || *p.WebPartType != *input.WebPartType) {
		return false
	}

	return true
}
