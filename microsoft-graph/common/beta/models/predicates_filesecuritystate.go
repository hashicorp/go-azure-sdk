package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileSecurityStateOperationPredicate struct {
	Name      *string
	ODataType *string
	Path      *string
	RiskScore *string
}

func (p FileSecurityStateOperationPredicate) Matches(input FileSecurityState) bool {

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Path != nil && (input.Path == nil || *p.Path != *input.Path) {
		return false
	}

	if p.RiskScore != nil && (input.RiskScore == nil || *p.RiskScore != *input.RiskScore) {
		return false
	}

	return true
}
