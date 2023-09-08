package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPartnerAgentInstallResultOperationPredicate struct {
	IsThirdPartyPartner *bool
	ODataType           *string
	Retriable           *bool
}

func (p CloudPCPartnerAgentInstallResultOperationPredicate) Matches(input CloudPCPartnerAgentInstallResult) bool {

	if p.IsThirdPartyPartner != nil && (input.IsThirdPartyPartner == nil || *p.IsThirdPartyPartner != *input.IsThirdPartyPartner) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Retriable != nil && (input.Retriable == nil || *p.Retriable != *input.Retriable) {
		return false
	}

	return true
}
