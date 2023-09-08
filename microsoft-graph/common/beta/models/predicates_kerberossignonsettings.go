package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KerberosSignOnSettingsOperationPredicate struct {
	KerberosServicePrincipalName *string
	ODataType                    *string
}

func (p KerberosSignOnSettingsOperationPredicate) Matches(input KerberosSignOnSettings) bool {

	if p.KerberosServicePrincipalName != nil && (input.KerberosServicePrincipalName == nil || *p.KerberosServicePrincipalName != *input.KerberosServicePrincipalName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
