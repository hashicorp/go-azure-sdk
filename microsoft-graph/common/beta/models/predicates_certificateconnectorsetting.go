package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateConnectorSettingOperationPredicate struct {
	CertExpiryTime              *string
	ConnectorVersion            *string
	EnrollmentError             *string
	LastConnectorConnectionTime *string
	LastUploadVersion           *int64
	ODataType                   *string
	Status                      *int64
}

func (p CertificateConnectorSettingOperationPredicate) Matches(input CertificateConnectorSetting) bool {

	if p.CertExpiryTime != nil && (input.CertExpiryTime == nil || *p.CertExpiryTime != *input.CertExpiryTime) {
		return false
	}

	if p.ConnectorVersion != nil && (input.ConnectorVersion == nil || *p.ConnectorVersion != *input.ConnectorVersion) {
		return false
	}

	if p.EnrollmentError != nil && (input.EnrollmentError == nil || *p.EnrollmentError != *input.EnrollmentError) {
		return false
	}

	if p.LastConnectorConnectionTime != nil && (input.LastConnectorConnectionTime == nil || *p.LastConnectorConnectionTime != *input.LastConnectorConnectionTime) {
		return false
	}

	if p.LastUploadVersion != nil && (input.LastUploadVersion == nil || *p.LastUploadVersion != *input.LastUploadVersion) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}
