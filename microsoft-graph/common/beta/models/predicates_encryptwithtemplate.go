package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptWithTemplateOperationPredicate struct {
	AvailableForEncryption *bool
	Name                   *string
	ODataType              *string
	TemplateId             *string
}

func (p EncryptWithTemplateOperationPredicate) Matches(input EncryptWithTemplate) bool {

	if p.AvailableForEncryption != nil && (input.AvailableForEncryption == nil || *p.AvailableForEncryption != *input.AvailableForEncryption) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TemplateId != nil && (input.TemplateId == nil || *p.TemplateId != *input.TemplateId) {
		return false
	}

	return true
}
