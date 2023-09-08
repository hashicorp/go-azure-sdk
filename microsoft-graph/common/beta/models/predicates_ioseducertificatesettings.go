package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEduCertificateSettingsOperationPredicate struct {
	CertFileName                   *string
	CertificateTemplateName        *string
	CertificateValidityPeriodValue *int64
	CertificationAuthority         *string
	CertificationAuthorityName     *string
	ODataType                      *string
	RenewalThresholdPercentage     *int64
	TrustedRootCertificate         *string
}

func (p IosEduCertificateSettingsOperationPredicate) Matches(input IosEduCertificateSettings) bool {

	if p.CertFileName != nil && (input.CertFileName == nil || *p.CertFileName != *input.CertFileName) {
		return false
	}

	if p.CertificateTemplateName != nil && (input.CertificateTemplateName == nil || *p.CertificateTemplateName != *input.CertificateTemplateName) {
		return false
	}

	if p.CertificateValidityPeriodValue != nil && (input.CertificateValidityPeriodValue == nil || *p.CertificateValidityPeriodValue != *input.CertificateValidityPeriodValue) {
		return false
	}

	if p.CertificationAuthority != nil && (input.CertificationAuthority == nil || *p.CertificationAuthority != *input.CertificationAuthority) {
		return false
	}

	if p.CertificationAuthorityName != nil && (input.CertificationAuthorityName == nil || *p.CertificationAuthorityName != *input.CertificationAuthorityName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RenewalThresholdPercentage != nil && (input.RenewalThresholdPercentage == nil || *p.RenewalThresholdPercentage != *input.RenewalThresholdPercentage) {
		return false
	}

	if p.TrustedRootCertificate != nil && (input.TrustedRootCertificate == nil || *p.TrustedRootCertificate != *input.TrustedRootCertificate) {
		return false
	}

	return true
}
