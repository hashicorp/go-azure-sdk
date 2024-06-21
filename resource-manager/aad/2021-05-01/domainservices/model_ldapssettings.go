package domainservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LdapsSettings struct {
	CertificateNotAfter    *string         `json:"certificateNotAfter,omitempty"`
	CertificateThumbprint  *string         `json:"certificateThumbprint,omitempty"`
	ExternalAccess         *ExternalAccess `json:"externalAccess,omitempty"`
	Ldaps                  *Ldaps          `json:"ldaps,omitempty"`
	PfxCertificate         *string         `json:"pfxCertificate,omitempty"`
	PfxCertificatePassword *string         `json:"pfxCertificatePassword,omitempty"`
	PublicCertificate      *string         `json:"publicCertificate,omitempty"`
}
