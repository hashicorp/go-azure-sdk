package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationTemplateOperationPredicate struct {
	ApplicationId *string
	Default       *bool
	Description   *string
	Discoverable  *bool
	FactoryTag    *string
	Id            *string
	ODataType     *string
}

func (p SynchronizationTemplateOperationPredicate) Matches(input SynchronizationTemplate) bool {

	if p.ApplicationId != nil && (input.ApplicationId == nil || *p.ApplicationId != *input.ApplicationId) {
		return false
	}

	if p.Default != nil && (input.Default == nil || *p.Default != *input.Default) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Discoverable != nil && (input.Discoverable == nil || *p.Discoverable != *input.Discoverable) {
		return false
	}

	if p.FactoryTag != nil && (input.FactoryTag == nil || *p.FactoryTag != *input.FactoryTag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
