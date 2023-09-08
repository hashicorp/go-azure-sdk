package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceInformationOperationPredicate struct {
	CertificationName *string
	ODataType         *string
}

func (p ComplianceInformationOperationPredicate) Matches(input ComplianceInformation) bool {

	if p.CertificationName != nil && (input.CertificationName == nil || *p.CertificationName != *input.CertificationName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
