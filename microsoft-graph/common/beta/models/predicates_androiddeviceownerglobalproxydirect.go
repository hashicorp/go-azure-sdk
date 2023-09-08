package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGlobalProxyDirectOperationPredicate struct {
	Host      *string
	ODataType *string
	Port      *int64
}

func (p AndroidDeviceOwnerGlobalProxyDirectOperationPredicate) Matches(input AndroidDeviceOwnerGlobalProxyDirect) bool {

	if p.Host != nil && (input.Host == nil || *p.Host != *input.Host) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Port != nil && (input.Port == nil || *p.Port != *input.Port) {
		return false
	}

	return true
}
