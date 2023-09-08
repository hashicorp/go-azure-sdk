package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessLocalConnectivityConfigurationOperationPredicate struct {
	Asn        *int64
	BgpAddress *string
	Endpoint   *string
	ODataType  *string
}

func (p NetworkaccessLocalConnectivityConfigurationOperationPredicate) Matches(input NetworkaccessLocalConnectivityConfiguration) bool {

	if p.Asn != nil && (input.Asn == nil || *p.Asn != *input.Asn) {
		return false
	}

	if p.BgpAddress != nil && (input.BgpAddress == nil || *p.BgpAddress != *input.BgpAddress) {
		return false
	}

	if p.Endpoint != nil && (input.Endpoint == nil || *p.Endpoint != *input.Endpoint) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
