package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerSilentCertificateAccessOperationPredicate struct {
	ODataType *string
	PackageId *string
}

func (p AndroidDeviceOwnerSilentCertificateAccessOperationPredicate) Matches(input AndroidDeviceOwnerSilentCertificateAccess) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PackageId != nil && (input.PackageId == nil || *p.PackageId != *input.PackageId) {
		return false
	}

	return true
}
