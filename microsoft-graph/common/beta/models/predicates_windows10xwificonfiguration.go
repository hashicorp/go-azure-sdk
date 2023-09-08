package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XWifiConfigurationOperationPredicate struct {
	AuthenticationCertificateId *string
	CreationDateTime            *string
	CustomXml                   *string
	CustomXmlFileName           *string
	Description                 *string
	DisplayName                 *string
	Id                          *string
	LastModifiedDateTime        *string
	ODataType                   *string
	Version                     *int64
}

func (p Windows10XWifiConfigurationOperationPredicate) Matches(input Windows10XWifiConfiguration) bool {

	if p.AuthenticationCertificateId != nil && (input.AuthenticationCertificateId == nil || *p.AuthenticationCertificateId != *input.AuthenticationCertificateId) {
		return false
	}

	if p.CreationDateTime != nil && (input.CreationDateTime == nil || *p.CreationDateTime != *input.CreationDateTime) {
		return false
	}

	if p.CustomXml != nil && (input.CustomXml == nil || *p.CustomXml != *input.CustomXml) {
		return false
	}

	if p.CustomXmlFileName != nil && (input.CustomXmlFileName == nil || *p.CustomXmlFileName != *input.CustomXmlFileName) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
