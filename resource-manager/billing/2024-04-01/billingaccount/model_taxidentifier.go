package billingaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TaxIdentifier struct {
	Country *string              `json:"country,omitempty"`
	Id      *string              `json:"id,omitempty"`
	Scope   *string              `json:"scope,omitempty"`
	Status  *TaxIdentifierStatus `json:"status,omitempty"`
	Type    *TaxIdentifierType   `json:"type,omitempty"`
}
