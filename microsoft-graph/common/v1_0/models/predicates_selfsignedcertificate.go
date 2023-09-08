package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SelfSignedCertificateOperationPredicate struct {
	CustomKeyIdentifier *string
	DisplayName         *string
	EndDateTime         *string
	Key                 *string
	KeyId               *string
	ODataType           *string
	StartDateTime       *string
	Thumbprint          *string
	Type                *string
	Usage               *string
}

func (p SelfSignedCertificateOperationPredicate) Matches(input SelfSignedCertificate) bool {

	if p.CustomKeyIdentifier != nil && (input.CustomKeyIdentifier == nil || *p.CustomKeyIdentifier != *input.CustomKeyIdentifier) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.Key != nil && (input.Key == nil || *p.Key != *input.Key) {
		return false
	}

	if p.KeyId != nil && (input.KeyId == nil || *p.KeyId != *input.KeyId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.Thumbprint != nil && (input.Thumbprint == nil || *p.Thumbprint != *input.Thumbprint) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	if p.Usage != nil && (input.Usage == nil || *p.Usage != *input.Usage) {
		return false
	}

	return true
}
