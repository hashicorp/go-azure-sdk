package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSImportedPFXCertificateProfileIntendedPurpose string

const (
	MacOSImportedPFXCertificateProfileIntendedPurposesmimeEncryption MacOSImportedPFXCertificateProfileIntendedPurpose = "SmimeEncryption"
	MacOSImportedPFXCertificateProfileIntendedPurposesmimeSigning    MacOSImportedPFXCertificateProfileIntendedPurpose = "SmimeSigning"
	MacOSImportedPFXCertificateProfileIntendedPurposeunassigned      MacOSImportedPFXCertificateProfileIntendedPurpose = "Unassigned"
	MacOSImportedPFXCertificateProfileIntendedPurposevpn             MacOSImportedPFXCertificateProfileIntendedPurpose = "Vpn"
	MacOSImportedPFXCertificateProfileIntendedPurposewifi            MacOSImportedPFXCertificateProfileIntendedPurpose = "Wifi"
)

func PossibleValuesForMacOSImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(MacOSImportedPFXCertificateProfileIntendedPurposesmimeEncryption),
		string(MacOSImportedPFXCertificateProfileIntendedPurposesmimeSigning),
		string(MacOSImportedPFXCertificateProfileIntendedPurposeunassigned),
		string(MacOSImportedPFXCertificateProfileIntendedPurposevpn),
		string(MacOSImportedPFXCertificateProfileIntendedPurposewifi),
	}
}

func parseMacOSImportedPFXCertificateProfileIntendedPurpose(input string) (*MacOSImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]MacOSImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": MacOSImportedPFXCertificateProfileIntendedPurposesmimeEncryption,
		"smimesigning":    MacOSImportedPFXCertificateProfileIntendedPurposesmimeSigning,
		"unassigned":      MacOSImportedPFXCertificateProfileIntendedPurposeunassigned,
		"vpn":             MacOSImportedPFXCertificateProfileIntendedPurposevpn,
		"wifi":            MacOSImportedPFXCertificateProfileIntendedPurposewifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
