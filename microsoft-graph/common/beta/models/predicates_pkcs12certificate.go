package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Pkcs12CertificateOperationPredicate struct {
	ODataType   *string
	Password    *string
	Pkcs12Value *string
}

func (p Pkcs12CertificateOperationPredicate) Matches(input Pkcs12Certificate) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Password != nil && (input.Password == nil || *p.Password != *input.Password) {
		return false
	}

	if p.Pkcs12Value != nil && (input.Pkcs12Value == nil || *p.Pkcs12Value != *input.Pkcs12Value) {
		return false
	}

	return true
}
