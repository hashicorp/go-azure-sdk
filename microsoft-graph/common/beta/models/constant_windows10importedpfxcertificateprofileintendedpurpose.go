package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10ImportedPFXCertificateProfileIntendedPurpose string

const (
	Windows10ImportedPFXCertificateProfileIntendedPurposesmimeEncryption Windows10ImportedPFXCertificateProfileIntendedPurpose = "SmimeEncryption"
	Windows10ImportedPFXCertificateProfileIntendedPurposesmimeSigning    Windows10ImportedPFXCertificateProfileIntendedPurpose = "SmimeSigning"
	Windows10ImportedPFXCertificateProfileIntendedPurposeunassigned      Windows10ImportedPFXCertificateProfileIntendedPurpose = "Unassigned"
	Windows10ImportedPFXCertificateProfileIntendedPurposevpn             Windows10ImportedPFXCertificateProfileIntendedPurpose = "Vpn"
	Windows10ImportedPFXCertificateProfileIntendedPurposewifi            Windows10ImportedPFXCertificateProfileIntendedPurpose = "Wifi"
)

func PossibleValuesForWindows10ImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(Windows10ImportedPFXCertificateProfileIntendedPurposesmimeEncryption),
		string(Windows10ImportedPFXCertificateProfileIntendedPurposesmimeSigning),
		string(Windows10ImportedPFXCertificateProfileIntendedPurposeunassigned),
		string(Windows10ImportedPFXCertificateProfileIntendedPurposevpn),
		string(Windows10ImportedPFXCertificateProfileIntendedPurposewifi),
	}
}

func parseWindows10ImportedPFXCertificateProfileIntendedPurpose(input string) (*Windows10ImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]Windows10ImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": Windows10ImportedPFXCertificateProfileIntendedPurposesmimeEncryption,
		"smimesigning":    Windows10ImportedPFXCertificateProfileIntendedPurposesmimeSigning,
		"unassigned":      Windows10ImportedPFXCertificateProfileIntendedPurposeunassigned,
		"vpn":             Windows10ImportedPFXCertificateProfileIntendedPurposevpn,
		"wifi":            Windows10ImportedPFXCertificateProfileIntendedPurposewifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10ImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
