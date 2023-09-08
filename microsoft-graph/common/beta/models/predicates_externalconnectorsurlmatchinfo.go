package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsUrlMatchInfoOperationPredicate struct {
	ODataType  *string
	UrlPattern *string
}

func (p ExternalConnectorsUrlMatchInfoOperationPredicate) Matches(input ExternalConnectorsUrlMatchInfo) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UrlPattern != nil && (input.UrlPattern == nil || *p.UrlPattern != *input.UrlPattern) {
		return false
	}

	return true
}
