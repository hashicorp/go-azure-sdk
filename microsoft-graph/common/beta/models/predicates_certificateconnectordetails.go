package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateConnectorDetailsOperationPredicate struct {
	ConnectorName       *string
	ConnectorVersion    *string
	EnrollmentDateTime  *string
	Id                  *string
	LastCheckinDateTime *string
	MachineName         *string
	ODataType           *string
}

func (p CertificateConnectorDetailsOperationPredicate) Matches(input CertificateConnectorDetails) bool {

	if p.ConnectorName != nil && (input.ConnectorName == nil || *p.ConnectorName != *input.ConnectorName) {
		return false
	}

	if p.ConnectorVersion != nil && (input.ConnectorVersion == nil || *p.ConnectorVersion != *input.ConnectorVersion) {
		return false
	}

	if p.EnrollmentDateTime != nil && (input.EnrollmentDateTime == nil || *p.EnrollmentDateTime != *input.EnrollmentDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastCheckinDateTime != nil && (input.LastCheckinDateTime == nil || *p.LastCheckinDateTime != *input.LastCheckinDateTime) {
		return false
	}

	if p.MachineName != nil && (input.MachineName == nil || *p.MachineName != *input.MachineName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
