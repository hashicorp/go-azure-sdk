package certificate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateFormat string

const (
	CertificateFormatCer CertificateFormat = "Cer"
	CertificateFormatPfx CertificateFormat = "Pfx"
)

func PossibleValuesForCertificateFormat() []string {
	return []string{
		string(CertificateFormatCer),
		string(CertificateFormatPfx),
	}
}

type CertificateProvisioningState string

const (
	CertificateProvisioningStateDeleting  CertificateProvisioningState = "Deleting"
	CertificateProvisioningStateFailed    CertificateProvisioningState = "Failed"
	CertificateProvisioningStateSucceeded CertificateProvisioningState = "Succeeded"
)

func PossibleValuesForCertificateProvisioningState() []string {
	return []string{
		string(CertificateProvisioningStateDeleting),
		string(CertificateProvisioningStateFailed),
		string(CertificateProvisioningStateSucceeded),
	}
}
