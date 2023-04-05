package dpscertificate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificatePurpose string

const (
	CertificatePurposeClientAuthentication CertificatePurpose = "clientAuthentication"
	CertificatePurposeServerAuthentication CertificatePurpose = "serverAuthentication"
)

func PossibleValuesForCertificatePurpose() []string {
	return []string{
		string(CertificatePurposeClientAuthentication),
		string(CertificatePurposeServerAuthentication),
	}
}
