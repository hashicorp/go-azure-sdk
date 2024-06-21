package cacertificates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CaCertificateProperties struct {
	Description        *string                         `json:"description,omitempty"`
	EncodedCertificate *string                         `json:"encodedCertificate,omitempty"`
	ExpiryTimeInUtc    *string                         `json:"expiryTimeInUtc,omitempty"`
	IssueTimeInUtc     *string                         `json:"issueTimeInUtc,omitempty"`
	ProvisioningState  *CaCertificateProvisioningState `json:"provisioningState,omitempty"`
}
