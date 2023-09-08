package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose string

const (
	WindowsPhone81ImportedPFXCertificateProfileIntendedPurposesmimeEncryption WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose = "SmimeEncryption"
	WindowsPhone81ImportedPFXCertificateProfileIntendedPurposesmimeSigning    WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose = "SmimeSigning"
	WindowsPhone81ImportedPFXCertificateProfileIntendedPurposeunassigned      WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose = "Unassigned"
	WindowsPhone81ImportedPFXCertificateProfileIntendedPurposevpn             WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose = "Vpn"
	WindowsPhone81ImportedPFXCertificateProfileIntendedPurposewifi            WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose = "Wifi"
)

func PossibleValuesForWindowsPhone81ImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(WindowsPhone81ImportedPFXCertificateProfileIntendedPurposesmimeEncryption),
		string(WindowsPhone81ImportedPFXCertificateProfileIntendedPurposesmimeSigning),
		string(WindowsPhone81ImportedPFXCertificateProfileIntendedPurposeunassigned),
		string(WindowsPhone81ImportedPFXCertificateProfileIntendedPurposevpn),
		string(WindowsPhone81ImportedPFXCertificateProfileIntendedPurposewifi),
	}
}

func parseWindowsPhone81ImportedPFXCertificateProfileIntendedPurpose(input string) (*WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": WindowsPhone81ImportedPFXCertificateProfileIntendedPurposesmimeEncryption,
		"smimesigning":    WindowsPhone81ImportedPFXCertificateProfileIntendedPurposesmimeSigning,
		"unassigned":      WindowsPhone81ImportedPFXCertificateProfileIntendedPurposeunassigned,
		"vpn":             WindowsPhone81ImportedPFXCertificateProfileIntendedPurposevpn,
		"wifi":            WindowsPhone81ImportedPFXCertificateProfileIntendedPurposewifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
