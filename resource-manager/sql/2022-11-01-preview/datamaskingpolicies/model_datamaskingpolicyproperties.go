package datamaskingpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMaskingPolicyProperties struct {
	ApplicationPrincipals *string          `json:"applicationPrincipals,omitempty"`
	DataMaskingState      DataMaskingState `json:"dataMaskingState"`
	ExemptPrincipals      *string          `json:"exemptPrincipals,omitempty"`
	MaskingLevel          *string          `json:"maskingLevel,omitempty"`
}
