package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectByTemplateActionOperationPredicate struct {
	ODataType  *string
	TemplateId *string
}

func (p ProtectByTemplateActionOperationPredicate) Matches(input ProtectByTemplateAction) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TemplateId != nil && (input.TemplateId == nil || *p.TemplateId != *input.TemplateId) {
		return false
	}

	return true
}
