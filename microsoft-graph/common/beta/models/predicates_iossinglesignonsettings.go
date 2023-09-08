package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosSingleSignOnSettingsOperationPredicate struct {
	DisplayName           *string
	KerberosPrincipalName *string
	KerberosRealm         *string
	ODataType             *string
}

func (p IosSingleSignOnSettingsOperationPredicate) Matches(input IosSingleSignOnSettings) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.KerberosPrincipalName != nil && (input.KerberosPrincipalName == nil || *p.KerberosPrincipalName != *input.KerberosPrincipalName) {
		return false
	}

	if p.KerberosRealm != nil && (input.KerberosRealm == nil || *p.KerberosRealm != *input.KerberosRealm) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
