package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnSecurityAssociationParametersOperationPredicate struct {
	LifetimeInMinutes          *int64
	ODataType                  *string
	SecurityDiffieHellmanGroup *int64
}

func (p IosVpnSecurityAssociationParametersOperationPredicate) Matches(input IosVpnSecurityAssociationParameters) bool {

	if p.LifetimeInMinutes != nil && (input.LifetimeInMinutes == nil || *p.LifetimeInMinutes != *input.LifetimeInMinutes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SecurityDiffieHellmanGroup != nil && (input.SecurityDiffieHellmanGroup == nil || *p.SecurityDiffieHellmanGroup != *input.SecurityDiffieHellmanGroup) {
		return false
	}

	return true
}
