package agreements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementProperties struct {
	Accepted          *bool   `json:"accepted,omitempty"`
	LicenseTextLink   *string `json:"licenseTextLink,omitempty"`
	Plan              *string `json:"plan,omitempty"`
	PrivacyPolicyLink *string `json:"privacyPolicyLink,omitempty"`
	Product           *string `json:"product,omitempty"`
	Publisher         *string `json:"publisher,omitempty"`
	RetrieveDatetime  *string `json:"retrieveDatetime,omitempty"`
	Signature         *string `json:"signature,omitempty"`
}
