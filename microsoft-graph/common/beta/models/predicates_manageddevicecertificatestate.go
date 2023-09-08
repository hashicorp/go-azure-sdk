package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateOperationPredicate struct {
	CertificateEnhancedKeyUsage                   *string
	CertificateErrorCode                          *int64
	CertificateExpirationDateTime                 *string
	CertificateIssuanceDateTime                   *string
	CertificateIssuer                             *string
	CertificateKeyLength                          *int64
	CertificateLastIssuanceStateChangedDateTime   *string
	CertificateProfileDisplayName                 *string
	CertificateSerialNumber                       *string
	CertificateSubjectAlternativeNameFormatString *string
	CertificateSubjectNameFormatString            *string
	CertificateThumbprint                         *string
	CertificateValidityPeriod                     *int64
	DeviceDisplayName                             *string
	Id                                            *string
	LastCertificateStateChangeDateTime            *string
	ODataType                                     *string
	UserDisplayName                               *string
}

func (p ManagedDeviceCertificateStateOperationPredicate) Matches(input ManagedDeviceCertificateState) bool {

	if p.CertificateEnhancedKeyUsage != nil && (input.CertificateEnhancedKeyUsage == nil || *p.CertificateEnhancedKeyUsage != *input.CertificateEnhancedKeyUsage) {
		return false
	}

	if p.CertificateErrorCode != nil && (input.CertificateErrorCode == nil || *p.CertificateErrorCode != *input.CertificateErrorCode) {
		return false
	}

	if p.CertificateExpirationDateTime != nil && (input.CertificateExpirationDateTime == nil || *p.CertificateExpirationDateTime != *input.CertificateExpirationDateTime) {
		return false
	}

	if p.CertificateIssuanceDateTime != nil && (input.CertificateIssuanceDateTime == nil || *p.CertificateIssuanceDateTime != *input.CertificateIssuanceDateTime) {
		return false
	}

	if p.CertificateIssuer != nil && (input.CertificateIssuer == nil || *p.CertificateIssuer != *input.CertificateIssuer) {
		return false
	}

	if p.CertificateKeyLength != nil && (input.CertificateKeyLength == nil || *p.CertificateKeyLength != *input.CertificateKeyLength) {
		return false
	}

	if p.CertificateLastIssuanceStateChangedDateTime != nil && (input.CertificateLastIssuanceStateChangedDateTime == nil || *p.CertificateLastIssuanceStateChangedDateTime != *input.CertificateLastIssuanceStateChangedDateTime) {
		return false
	}

	if p.CertificateProfileDisplayName != nil && (input.CertificateProfileDisplayName == nil || *p.CertificateProfileDisplayName != *input.CertificateProfileDisplayName) {
		return false
	}

	if p.CertificateSerialNumber != nil && (input.CertificateSerialNumber == nil || *p.CertificateSerialNumber != *input.CertificateSerialNumber) {
		return false
	}

	if p.CertificateSubjectAlternativeNameFormatString != nil && (input.CertificateSubjectAlternativeNameFormatString == nil || *p.CertificateSubjectAlternativeNameFormatString != *input.CertificateSubjectAlternativeNameFormatString) {
		return false
	}

	if p.CertificateSubjectNameFormatString != nil && (input.CertificateSubjectNameFormatString == nil || *p.CertificateSubjectNameFormatString != *input.CertificateSubjectNameFormatString) {
		return false
	}

	if p.CertificateThumbprint != nil && (input.CertificateThumbprint == nil || *p.CertificateThumbprint != *input.CertificateThumbprint) {
		return false
	}

	if p.CertificateValidityPeriod != nil && (input.CertificateValidityPeriod == nil || *p.CertificateValidityPeriod != *input.CertificateValidityPeriod) {
		return false
	}

	if p.DeviceDisplayName != nil && (input.DeviceDisplayName == nil || *p.DeviceDisplayName != *input.DeviceDisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastCertificateStateChangeDateTime != nil && (input.LastCertificateStateChangeDateTime == nil || *p.LastCertificateStateChangeDateTime != *input.LastCertificateStateChangeDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserDisplayName != nil && (input.UserDisplayName == nil || *p.UserDisplayName != *input.UserDisplayName) {
		return false
	}

	return true
}
